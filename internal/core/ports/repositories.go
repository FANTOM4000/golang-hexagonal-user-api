package ports

import (
	"app/internal/core/domains"
	"context"
	"time"
)

type UserRepository interface {
	CreateUser(ctx context.Context,user domains.UserProperty)(domains.User,error)
	UpdateUserById(ctx context.Context,id string,user domains.UserMetadata)(error)
	DeleteUserById(ctx context.Context,id string)(error)
	GetUserById(ctx context.Context,id string)(domains.User,error)
	GetUserByUsername(ctx context.Context,username string)(domains.User,error)
}

type CacheRepository interface {
	Get(context context.Context,key string) (string, error)
	Set(context context.Context, key string, val string,expire time.Duration)error
	Delete(ctx context.Context,key string)error
}


