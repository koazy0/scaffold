package config

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"moyu/api/config/v1"
)

func (c *ControllerV1) GetExpense(ctx context.Context, req *v1.GetExpenseReq) (res *v1.GetExpenseRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
