// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package config

import (
	"context"

	"moyu/api/config/v1"
)

type IConfigV1 interface {
	UpdateConfig(ctx context.Context, req *v1.UpdateConfigReq) (res *v1.UpdateConfigRes, err error)
	GetConfig(ctx context.Context, req *v1.GetConfigReq) (res *v1.GetConfigRes, err error)
	UpdateExpense(ctx context.Context, req *v1.UpdateExpenseReq) (res *v1.UpdateExpenseRes, err error)
	GetExpense(ctx context.Context, req *v1.GetExpenseReq) (res *v1.GetExpenseRes, err error)
	UpdateWorkOvertime(ctx context.Context, req *v1.UpdateWorkOvertimeReq) (res *v1.UpdateWorkOvertimeRes, err error)
	GetWorkOvertime(ctx context.Context, req *v1.GetWorkOvertimeReq) (res *v1.GetWorkOvertimeRes, err error)
}
