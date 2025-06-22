package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"scaffold/internal/model"
)

type SignInReq struct {
	g.Meta `path:"/user/signin" method:"post" tags:"用户管理" summary:"用户已有账号登录"`
	model.UserSignIn
}
type SignInRes struct {
	model.UserSignInReply
}

type SignUpReq struct {
	g.Meta `path:"/user/signup" method:"post" tags:"用户管理" summary:"用户注册账号"`
	model.UserSignUp
}
type SignUpRes struct {
	model.UserSignUpReply
}

type EmailVerifyReq struct {
	g.Meta `path:"/emailverify" method:"post" tags:"用户管理" summary:"邮箱验证"`
	Email  string `json:"email" v:"required"`
	Token  string `json:"token" v:"required"` //一次性生成的验证码，先留着
}
type EmailVerifyRes struct{}
