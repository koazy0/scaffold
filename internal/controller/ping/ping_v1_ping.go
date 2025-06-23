package ping

import (
	"context"

	"scaffold/api/ping/v1"
)

func (c *ControllerV1) Ping(ctx context.Context, req *v1.PingReq) (res *v1.PingRes, err error) {
	return &v1.PingRes{
		Message: "pong",
	}, nil
}
