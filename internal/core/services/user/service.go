package userservice

import (
	"app/internal/core/domains"
	"app/internal/core/ports"
	"app/pkg/hashpass"
	jwtPkg "app/pkg/jwt"
	"app/pkg/standard"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type userService struct {
	userRepo ports.UserRepository
	cache    ports.CacheRepository
}

func NewUserService(userRepo ports.UserRepository, cache ports.CacheRepository) ports.UserService {
	return userService{userRepo: userRepo, cache: cache}
}

func (us userService) CreateUser(ctx context.Context, user domains.UserProperty) (domains.User, error) {
	hashPass, err := hashpass.HashPassword(user.Password)
	if err != nil {
		return domains.User{}, err
	}
	user.Password = hashPass
	user.CreatedAt = standard.Now()
	user.UpdatedAt = standard.Now()
	userObj, err := us.userRepo.CreateUser(ctx, user)
	userObj.Password = ""
	return userObj, err
}
func (us userService) UpdateUserById(ctx context.Context, id string, user domains.UserMetadata) error {
	if _,err := us.cache.Get(ctx, fmt.Sprintf(standard.AuthKey, id));err == redis.Nil{
		return err
	}
	err := us.userRepo.UpdateUserById(ctx, id, user)
	return err
}
func (us userService) DeleteUserById(ctx context.Context, id string) error {
	if _,err := us.cache.Get(ctx, fmt.Sprintf(standard.AuthKey, id));err == redis.Nil{
		return err
	}
	err := us.userRepo.DeleteUserById(ctx, id)
	return err
}
func (us userService) GetUserById(ctx context.Context, id string) (domains.User, error) {
	if _,err := us.cache.Get(ctx, fmt.Sprintf(standard.AuthKey, id));err == redis.Nil{
		return domains.User{},err
	}
	userObj, err := us.userRepo.GetUserById(ctx, id)
	return userObj, err
}
func (us userService) Login(ctx context.Context, auth domains.UserNamePassword) (string, error) {
	user, err := us.userRepo.GetUserByUsername(ctx, auth.Username)
	if err != nil {
		return "", err
	}
	pass := hashpass.CheckPasswordHash(auth.Password, user.Password)
	if !pass {
		return "", errors.New("invalid password")
	}
	//generate token without expire
	token := jwtPkg.GenerateToken(user.ID.Hex(), user.Role, 0)
	us.cache.Set(ctx, fmt.Sprintf(standard.AuthKey, user.ID.Hex()), token, time.Second*180)
	return token, nil
}

func (us userService) LogOut(ctx context.Context, id string)error{
	return us.cache.Delete(ctx,fmt.Sprintf(standard.AuthKey, id))
}
