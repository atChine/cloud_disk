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

type UserFolderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderCreateLogic {
	return &UserFolderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFolderCreateLogic) UserFolderCreate(req *types.UserFolderCreateRequest, userIdentity string) (resp *types.UserFolderCreateReply, err error) {
	// todo: add your logic here and delete this line
	if userIdentity == "" {
		return nil, errors.New("userIdentity is empty")
	}
	//首先判断同级目录下是否存在同名文件夹
	count, err := l.svcCtx.Engine.Where("parent_id = ? AND name = ? AND user_identity = ?", req.ParentId, req.Name, userIdentity).Count(new(models.UserRepository))
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("文件夹已存在")
	}
	//创建文件夹
	data := models.UserRepository{
		Identity:     helper.UUID(),
		UserIdentity: userIdentity,
		ParentId:     req.ParentId,
		Name:         req.Name,
	}
	insert, err := l.svcCtx.Engine.Insert(&data)
	if err != nil {
		return nil, err
	}
	if insert != 1 {
		return nil, errors.New("创建文件夹失败")
	}
	return &types.UserFolderCreateReply{
		Identity: data.Identity,
	}, nil
}
