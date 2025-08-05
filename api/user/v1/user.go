package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"moyu/internal/model"
)

type LoginReq struct {
	g.Meta `path:"/login" method:"post" tags:"用户管理" summary:"用户已有账号登录"`
	model.UserSignIn
}
type LoginRes struct {
	model.UserSignInReply
}
