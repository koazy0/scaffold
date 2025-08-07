package config

import (
	"context"
	"errors"
	v1 "moyu/api/config/v1"
	"moyu/internal/common"
)

type (
	sUser struct{}
)

var (
	logger = common.Logs().Cat("config")
)

func init() {
}

func GetConfigs(ctx context.Context, in *v1.GetWorkOvertimeReq) (res *v1.GetWorkOvertimeRes, err error) {
	userID, exist := ctx.Value("user_id").(string)
	if !exist {
		logger.Warn("user_id not exist")
		return nil, errors.New("Internal Server Error")
	}
	println(userID)
	return
}
