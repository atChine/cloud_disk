package logic

import (
	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"cloud_disk/core/models"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadPrepareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadPrepareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadPrepareLogic {
	return &FileUploadPrepareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadPrepareLogic) FileUploadPrepare(req *types.FileUploadPrePareRequest) (resp *types.FileUploadPrePareReply, err error) {
	rp := new(models.RepositoryPool)
	//查询仓库是否存在此文件，用hash值作为查询条件
	has, err := l.svcCtx.Engine.Where("hash = ?", req.Md5).Get(rp)
	if err != nil {
		return nil, err
	}
	if has {
		//秒传成功
		resp = new(types.FileUploadPrePareReply)
		resp.Identity = rp.Identity
	} else {
		// TODO:获取该文件的UploadId，用来进行文件分片上传
	}
	return
}
