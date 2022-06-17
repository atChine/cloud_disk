package logic

import (
	"cloud_disk/core/helper"
	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"cloud_disk/core/models"
	"context"
	"errors"
	"log"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterReplay, err error) {
	// todo: add your logic here and delete this line
	//判断code是否一致
	code, err := l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if err != nil {
		return nil, errors.New("邮箱验证码还为空")
	}
	if code != req.Code{
		err = errors.New("验证码错误")
		return
	}
	//判断用户名是否存在
	count, err := l.svcCtx.Engine.Where("name = ?", req.Name).Count(new(models.UserBasic))
	if err != nil {
		return nil, err
	}
	if count > 0{
		err = errors.New("用户名已存在")
		return
	}
	//将注册的信息存入数据库
	user := &models.UserBasic{
		Identity: helper.UUID(),
		Name:     req.Name,
		Password: helper.Md5(req.Password),
		Email:    req.Email,
	}
	count, err = l.svcCtx.Engine.Insert(user)
	if err != nil {
		return nil, err
	}
	if count == 0{
		return nil,errors.New("注册失败")
	}
	if count > 0{
		log.Println("注册成功")
	}
	return
}
