package config

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"moyu/api/config/v1"
)

func (c *ControllerV1) GetWorkOvertime(ctx context.Context, req *v1.GetWorkOvertimeReq) (res *v1.GetWorkOvertimeRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
