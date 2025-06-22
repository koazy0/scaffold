// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"scaffold/internal/model"
)

type (
	IUser interface {
		ValidateUser(ctx context.Context, in model.UserSignIn) (out *model.UserSignInReply, err error)
		CreateUser(ctx context.Context, in model.UserSignUp) (out *model.UserSignUpReply, err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
