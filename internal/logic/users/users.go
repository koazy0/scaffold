package users

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"scaffold/internal/dao"
	"scaffold/internal/logic/utils"
	"scaffold/internal/model"
	"scaffold/internal/model/do"
	"scaffold/internal/service"
	"strconv"
)

type (
	sUser struct{}
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
		zap.S().Error(err.Error())
		return nil, err
	}
	//根据加密后对比
	passwordEncryptIn := service.Jwt().HashPassword(in.Password, *userData.Salt)
	if passwordEncryptIn != *userData.Password {
		return nil, errors.New("password is incorrect")
	}
	token, err := service.Jwt().GenerateToken(ctx, in.UserID)
	if err != nil {
		zap.S().Error(err.Error())
		return nil, err
	}
	out = &model.UserSignInReply{
		Token: token,
	}
	return out, nil
}

func (s *sUser) CreateUser(ctx context.Context, in model.UserSignUp) (out *model.UserSignUpReply, err error) {

	//查询数据库中有多少个角色
	users := []do.UserModels{}
	count := 1
	err = dao.UserModels.Ctx(ctx).ScanAndCount(users, &count, false)
	if err != nil {
		zap.S().Error(err.Error())
		return nil, err
	}
	//根据length来进行uid的生成
	uid := 100000000 + count + 1
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
	res, err := dao.UserModels.Ctx(ctx).Insert(userModel)
	if err != nil {
		zap.S().Error(err.Error())
		return nil, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		zap.S().Error(err.Error())
		return nil, err
	}
	if affected == 0 {
		err = errors.New("affected rows equal 0")
		zap.S().Error(err.Error())
	}
	id, err := res.LastInsertId()
	if err != nil {
		zap.S().Error(err.Error())
		return nil, err
	}
	zap.S().Infof("new user created,userUID:%d, userID:%s, ID:%d", uid, in.UserID, id)
	return
}
