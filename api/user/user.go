// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"scaffold/api/user/v1"
)

type IUserV1 interface {
	SignIn(ctx context.Context, req *v1.SignInReq) (res *v1.SignInRes, err error)
	SignUp(ctx context.Context, req *v1.SignUpReq) (res *v1.SignUpRes, err error)
	EmailVerify(ctx context.Context, req *v1.EmailVerifyReq) (res *v1.EmailVerifyRes, err error)
}
