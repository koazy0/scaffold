package ping

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"scaffold/api/ping/v2"
)

func (c *ControllerV2) Ping(ctx context.Context, req *v2.PingReq) (res *v2.PingRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
