// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package ping

import (
	"context"

	"scaffold/api/ping/v1"
	"scaffold/api/ping/v2"
)

type IPingV1 interface {
	Ping(ctx context.Context, req *v1.PingReq) (res *v1.PingRes, err error)
}

type IPingV2 interface {
	Ping(ctx context.Context, req *v2.PingReq) (res *v2.PingRes, err error)
}
