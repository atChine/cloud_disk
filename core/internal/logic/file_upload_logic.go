package logic

import (
	"cloud_disk/core/helper"
	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"cloud_disk/core/models"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadLogic) FileUpload(req *types.FileUploadRequest) (resp *types.FileUploadReply, err error) {
	// todo: add your logic here and delete this line
	pool := models.RepositoryPool{
		Identity:  helper.UUID(),
		Hash:      req.Hash,
		Name:      req.Name,
		Ext:       req.Ext,
		Size:      req.Size,
		Path:      req.Path,
	}
	_, err = l.svcCtx.Engine.Insert(&pool)
	if err != nil {
		return nil, err
	}
	resp = new(types.FileUploadReply)
	resp.Identity = pool.Identity
	resp.Name = pool.Name
	resp.Ext = pool.Ext
	return
}
