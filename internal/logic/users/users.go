package users

import (
	"context"
	"errors"
	"fmt"
	"moyu/internal/common"
	"moyu/internal/dao"
	"moyu/internal/model"
	"moyu/internal/service"
)

type (
	sUser struct{}
)

var (
	logger = common.Logs().Cat("users")
)

func init() {
	service.RegisterUser(User())
}

// todo 后面改成依赖注入模式
var insUser = sUser{}

func User() *sUser {
	return &insUser
}

func (s *sUser) ValidateUser(ctx context.Context, in model.UserSignIn) (out *model.UserSignInReply, err error) {

	// 先从数据库里面查用户
	userData := model.UserModel{}

	err = dao.UserModels.Ctx(ctx).Where("user_id=?", in.UserID).Scan(&userData)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	//根据加密后对比

	token, err := service.Jwt().GenerateToken(ctx, in.UserID)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	out = &model.UserSignInReply{
		Token: token,
	}
	return out, nil
}

func (s *sUser) InitUser(ctx context.Context, userID string) (err error) {

	//查询数据库中有多少个角色
	//users := []do.UserModels{}
	//count := 1
	//err = dao.UserModels.Ctx(ctx).ScanAndCount(users, &count, false)
	//if err != nil {
	//	logger.Error(err.Error())
	//	return nil, err
	//}
	//根据length来进行uid的生成

	//创建用户
	userNoPasswdModel := model.UserNoPasswordModel{
		UserID: userID,
		//CreatedAt: time.Now(),
	}
	//存入usermodel
	res, err := dao.UserModels.Ctx(ctx).Insert(userNoPasswdModel)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	if affected == 0 {
		err = errors.New("affected rows equal 0")
		logger.Error(err.Error())
	}
	logger.Infof("new user created, userID:%s", userID)
	//创建默认配置
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	defaultConfig := model.UserConfigModel{
		UserID:        id,
		Income:        8000,
		IncomeCycle:   22,         //默认一个月上22天
		WorkTimeStart: "09:00:00", //time.TimeOnly
		WorkTimeEnd:   "18:00:00",
	}
	fmt.Println(defaultConfig)
	return nil
}
