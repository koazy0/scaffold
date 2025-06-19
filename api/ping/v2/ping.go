package v2

import (
	"github.com/gogf/gf/v2/frame/g"
)

type PingReq struct {
	g.Meta `path:"/ping" tags:"ping" method:"get" summary:"测试搭建成功"`
}
type PingRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
