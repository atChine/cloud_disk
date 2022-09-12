package logic

import (
	"cloud_disk/core/define"
	"cloud_disk/core/helper"
	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"cloud_disk/core/models"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginReplay, err error) {
	// todo: add your logic here and delete this line
	//从数据库中查询当前用户
	user := new(models.UserBasic)
	has, err := l.svcCtx.Engine.Where("name = ? AND password = ?", req.Name, helper.Md5(req.Password)).Get(user)
	if err != nil {
		return nil, errors.New("查询用户失败")
	}
	if !has {
		return nil, errors.New("用户名或者密码错误")
	}
	//生成token
	token, err := helper.GenerateToken(uint64(user.Id), user.Identity, user.Name, define.TokenExpire)
	if err != nil {
		return nil, err
	}
	//刷新token
	generateToken, err := helper.GenerateToken(uint64(user.Id), user.Identity, user.Name, define.RefreshTokenExpire)
	if err != nil {
		return nil, err
	}
	resp = new(types.LoginReplay)
	resp.Token = token
	resp.RefreshToken = generateToken
	return
}
