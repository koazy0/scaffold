package v1

import "github.com/gogf/gf/v2/frame/g"

type SignInReq struct {
	g.Meta   `path:"/user/signin" method:"post" tags:"用户管理" summary:"用户已有账号登录"`
	Username string `json:"username" v:"required"`
	Password string `json:"password" v:"required"`
}
type SignInRes struct{}

type SignUpReq struct {
	g.Meta `path:"/user/signup" method:"post" tags:"用户管理" summary:"用户注册账号"`
	//	UID  string `json:"username" v:"required"`	//ID还是后台自动生成吧
	UserID    string `json:"userid" v:"required"`
	Username  string `json:"username" v:"required"`
	Password  string `json:"password" v:"required"`
	Password2 string `json:"password2" v:"required"`
	Email     string `json:"email" v:"required"`
	//VerifyCode    string `json:"verify_code" v:"required"`//邮箱验证码，后面也许可能会用到
}
type SignUpRes struct{}

type EmailVerifyReq struct {
	g.Meta `path:"/emailverify" method:"post" tags:"用户管理" summary:"邮箱验证"`
	Email  string `json:"email" v:"required"`
	Token  string `json:"token" v:"required"` //一次性生成的验证码，先留着
}
type EmailVerifyRes struct{}
