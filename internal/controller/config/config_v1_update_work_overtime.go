package config

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"moyu/api/config/v1"
)

func (c *ControllerV1) UpdateWorkOvertime(ctx context.Context, req *v1.UpdateWorkOvertimeReq) (res *v1.UpdateWorkOvertimeRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
