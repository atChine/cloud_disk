package logic

import (
	"cloud_disk/core/helper"
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadChunkEndLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadChunkEndLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadChunkEndLogic {
	return &FileUploadChunkEndLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadChunkEndLogic) FileUploadChunkEnd(req *types.FileUploadChunkEndRequest) (resp *types.FileUploadChunkEndReply, err error) {

	co := make([]cos.Object, 0)
	for _, V := range req.CosObject {
		co = append(co, cos.Object{
			ETag:       V.ETag,
			PartNumber: V.PartNumber,
		})
	}
	err = helper.CosCompleteUpload(req.Key, req.UploadId, co)
	return
}
