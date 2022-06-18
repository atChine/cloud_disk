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

type ShareFileSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareFileSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareFileSaveLogic {
	return &ShareFileSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareFileSaveLogic) ShareFileSave(req *types.ShareFileSaveRequest, userIdentity string) (resp *types.ShareFileSaveReply, err error) {
	//获取资源
	rp := new(models.RepositoryPool)
	has, err := l.svcCtx.Engine.Where("identity = ?", req.RepositoryIdentity).Get(rp)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("资源不存在")
	}
	//保存分享资源
	uuid := helper.UUID()
	ur := &models.UserRepository{
		Identity:           uuid,
		UserIdentity:       userIdentity,
		ParentId:           req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
		Ext:                rp.Ext,
		Name:               rp.Name,
	}
	insert, err := l.svcCtx.Engine.Insert(ur)
	if err != nil {
		return nil, err
	}
	if insert != 1 {
		return nil, errors.New("保存分享资源失败")
	}
	return &types.ShareFileSaveReply{
		Identity: uuid,
	}, nil

}
