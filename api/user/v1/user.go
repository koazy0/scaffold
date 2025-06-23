package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"scaffold/internal/model"
)

type SignInReq struct {
	g.Meta `path:"/signin" method:"post" tags:"用户管理" summary:"用户已有账号登录"`
	model.UserSignIn
}
type SignInRes struct {
	model.UserSignInReply
}

type SignUpReq struct {
	g.Meta `path:"/signup" method:"post" tags:"用户管理" summary:"用户注册账号"`
	model.UserSignUp
}
type SignUpRes struct {
	model.UserSignUpReply
}

type LoginReq struct {
	g.Meta `path:"/login" method:"post" tags:"用户管理" summary:"用户已有账号登录"`
	model.UserSignIn
}
type LoginRes struct {
	model.UserSignInReply
}

// todo 再写这个路由的时候需要在头部加上JWT
type LogoutReq struct {
	g.Meta `path:"/logout" method:"get" tags:"用户管理" summary:"用户登出账号"`
	model.UserSignUp
}
type LogoutRes struct {
	model.UserSignUpReply
}

type EmailVerifyReq struct {
	g.Meta `path:"/emailverify" method:"post" tags:"用户管理" summary:"邮箱验证"`
	Email  string `json:"email" v:"required"`
	Token  string `json:"token" v:"required"` //一次性生成的验证码，先留着
}
type EmailVerifyRes struct{}
