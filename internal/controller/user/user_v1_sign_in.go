package user

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"scaffold/internal/model"
	"scaffold/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"

	"scaffold/api/user/v1"
)

func (c *ControllerV1) SignIn(ctx context.Context, req *v1.SignInReq) (res *v1.SignInRes, err error) {
	if req.UserID == "" || req.Password == "" {
		return nil, gerror.New("用户名或密码不能为空")
	}
	in := model.UserSignIn{}
	gconv.Struct(req, &in)
	out, err := service.User().ValidateUser(ctx, in)
	gconv.Struct(out, &res)
	return
}
