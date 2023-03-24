package ports

import (
	"app/internal/core/domains"
	"context"
)

type UserService interface {
	CreateUser(ctx context.Context,user domains.UserProperty)(domains.User,error)
	UpdateUserById(ctx context.Context,id string,user domains.UserMetadata)(error)
	DeleteUserById(ctx context.Context,id string)(error)
	GetUserById(ctx context.Context,id string)(domains.User,error)
	Login(ctx context.Context,auth domains.UserNamePassword)(string,error)
	LogOut(ctx context.Context,id string)(error)
}


