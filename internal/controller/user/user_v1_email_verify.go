package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"scaffold/api/user/v1"
)

func (c *ControllerV1) EmailVerify(ctx context.Context, req *v1.EmailVerifyReq) (res *v1.EmailVerifyRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
