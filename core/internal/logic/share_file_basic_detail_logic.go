package logic

import (
	"context"
	"errors"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareFileBasicDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareFileBasicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareFileBasicDetailLogic {
	return &ShareFileBasicDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareFileBasicDetailLogic) ShareFileBasicDetail(req *types.ShareFileBasicDetailRequest) (resp *types.ShareFileBasicDetailReplay, err error) {
	// 对分享记录的点击次数进行 + 1
	_, err = l.svcCtx.Engine.Exec("UPDATE share_basic SET click_num = click_num + 1 WHERE identity = ?", req.Identity)
	if err != nil {
		return
	}
	resp = new(types.ShareFileBasicDetailReplay)
	// 获取资源的详细信息
	get, err := l.svcCtx.Engine.Table("share_basic").
		Select("share_basic.repository_identity, user_repository.name, repository_pool.ext, repository_pool.size, repository_pool.path").
		Join("LEFT", "repository_pool", "share_basic.repository_identity = repository_pool.identity").
		Join("LEFT", "user_repository", "user_repository.identity = share_basic.user_repository_identity").
		Where("share_basic.identity = ?", req.Identity).Get(resp)
	if err != nil {
		return
	}
	if !get {
		return nil, errors.New("没有找到该分享记录")
	}

	return
}
