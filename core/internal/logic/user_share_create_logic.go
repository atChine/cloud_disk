package logic

import (
	"cloud_disk/core/helper"
	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"cloud_disk/core/models"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserShareCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserShareCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserShareCreateLogic {
	return &UserShareCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserShareCreateLogic) UserShareCreate(req *types.UserShareCreateRequest, userIdentity string) (resp *types.UserShareCreateReply, err error) {
	// todo: add your logic here and delete this line
	if userIdentity == "" {
		return nil, errors.New("用户身份不能为空")
	}
	uuid := helper.UUID()
	ur := new(models.UserRepository)
	get, err := l.svcCtx.Engine.Where("identity = ?", req.UserRepositoryIdentity).Get(ur)
	if err != nil {
		return nil, errors.New("没有identity，查询用户失败")
	}
	if !get {
		return nil, errors.New("没有identity，查询用户失败")
	}
	data := &models.ShareBasic{
		Identity:               uuid,
		UserIdentity:           userIdentity,
		UserRepositoryIdentity: req.UserRepositoryIdentity,
		RepositoryIdentity:     ur.RepositoryIdentity,
		ExpiredTime:            req.ExpiredTime,
	}
	insert, err := l.svcCtx.Engine.Insert(data)
	if err != nil {
		return nil, errors.New("创建分享失败")
	}
	if insert != 1 {
		return nil, errors.New("创建分享失败")
	}
	resp = new(types.UserShareCreateReply)
	resp.Identity = uuid
	return
}
