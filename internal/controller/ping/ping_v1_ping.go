package ping

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"scaffold/api/ping/v1"
)

func (c *ControllerV1) Ping(ctx context.Context, req *v1.PingReq) (res *v1.PingRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
