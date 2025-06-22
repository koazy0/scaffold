package users

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"scaffold/internal/dao"
	"scaffold/internal/logic/utils"
	"scaffold/internal/model"
	"scaffold/internal/service"
	"strconv"
)

type (
	sUser struct{}
)

func init() {
	service.RegisterUser(User())
}

var insUser = sUser{}

func User() *sUser {
	return &insUser
}

func (s *sUser) ValidateUser(ctx context.Context, in model.UserSignIn) (out *model.UserSignInReply, err error) {

	// 先从数据库里面查用户
	userData := model.UserModel{}

	//查不到,报错
	//根据加密后对比
	passwordEncryptIn := service.Jwt().HashPassword(in.Password, *userData.Salt)
	if passwordEncryptIn != *userData.Password {
		return &model.UserSignInReply{Message: "password is incorrect"}, nil
	}
	return
}

func (s *sUser) SignUpUser(ctx context.Context, in model.UserSignUp) (out *model.UserSignUpReply, err error) {

	// 先从数据库里面查用户
	userData := make([]model.UserModel, 0)

	//查询数据库中有多少个角色
	//根据length来进行uid的生成
	uid := 100000000 + len(userData)
	fmt.Println(dao.UserModels.Table())
	fmt.Println(dao.UserModels.Group())
	fmt.Println(dao.UserModels.Columns().UserId)
	fmt.Println(dao.UserModels.Columns().Salt)
	//fmt.Println(dao.UserModels.DB().)
	//dao.UserModels.Ctx(ctx).ScanList(&userData)
	salt, err := service.Jwt().GenerateSalt()
	if err != nil {
		zap.S().Errorln("生成盐值失败: " + err.Error())
		return nil, errors.New("生成盐值失败")
	}
	hashPassword := service.Jwt().HashPassword(in.Password, salt)
	userModel := model.UserModel{
		UID:      utils.CreatePointer(strconv.Itoa(uid)),
		UserID:   utils.CreatePointer(in.UserID),
		Username: utils.CreatePointer(in.Username),
		Salt:     utils.CreatePointer(salt),
		Password: utils.CreatePointer(hashPassword),
		Role:     2,
		Status:   1,
	}
	fmt.Println(userModel)
	//存入usermodel
	return
}
