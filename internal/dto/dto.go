package dto

import "app/internal/core/domains"

type BaseOKResponse struct {
	Code    int  `json:"code"`
	Message string `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type CreateUserRequest struct {
	domains.UserProperty
}

type UpdateUserRequest struct {
	domains.UserMetadata
}

type LoginRequest struct {
	domains.UserNamePassword
}

type LoginResponse struct {
	Token string `json:"token"`
}