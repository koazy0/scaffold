package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"moyu/internal/model"
)

// 增删改查配置
type UpdateConfigReq struct {
	g.Meta `path:"/config" method:"post" tags:"用户配置管理" summary:"用户更改配置"`
	model.UpdateConfigReq
}
type UpdateConfigRes struct {
	model.UpdateConfigReply
}

type GetConfigReq struct {
	g.Meta `path:"/config" method:"get" tags:"用户配置管理" summary:"用户获取配置"`
	model.GetConfigReq
}
type GetConfigRes struct {
	model.GetConfigReply
}

// 增删改查花销,只留一个统一的更改接口

type UpdateExpenseReq struct {
	g.Meta `path:"/expense" method:"post" tags:"用户花销管理" summary:"用户更改花销项"`
	model.UpdateExpenseReq
}
type UpdateExpenseRes struct {
	model.UpdateExpenseReply
}

type GetExpenseReq struct {
	g.Meta `path:"/expense" method:"get" tags:"用户花销管理" summary:"用户获取花销项"`
	model.GetExpenseReq
}
type GetExpenseRes struct {
	model.GetExpenseReply
}

type UpdateWorkOvertimeReq struct {
	g.Meta `path:"/overtime" method:"post" tags:"用户加班管理" summary:"用户更改用户加班配置"`
	model.UpdateWorkOverTimeReq
}
type UpdateWorkOvertimeRes struct {
	model.UpdateWorkOverTimeReply
}

type GetWorkOvertimeReq struct {
	g.Meta `path:"/overtime" method:"get" tags:"用户加班管理" summary:"用户获取用户加班配置"`
	model.GetWorkOverTimeReq
}
type GetWorkOvertimeRes struct {
	model.GetWorkOverTimeReply
}
