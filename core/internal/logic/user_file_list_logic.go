package logic

import (
	"cloud_disk/core/define"
	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"cloud_disk/core/models"
	"context"
	"errors"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest, userIdentity string) (resp *types.UserFileListReply, err error) {
	// todo: add your logic here and delete this line
	userFile := make([]*types.UserFile, 0)
	var count int64
	resp = new(types.UserFileListReply)
	//分页参数
	size := req.Size
	if size == 0 {
		size = define.PageSize
	}
	page := req.Page
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * size
	//查询用户文件列表
	err = l.svcCtx.Engine.Table("user_repository").Where("parent_id = ? AND user_identity = ? ", req.Id, userIdentity).
		Select("user_repository.id, user_repository.identity, user_repository.repository_identity, user_repository.ext,"+
			"user_repository.name, repository_pool.path, repository_pool.size").
		Join("LEFT", "repository_pool", "user_repository.repository_identity = repository_pool.identity").
		Where("user_repository.deleted_at = ? OR user_repository.deleted_at IS NULL", time.Time{}.Format(define.TimeFormat)).
		Limit(size, offset).Find(&userFile)
	if err != nil {
		return nil, errors.New("查询用户文件列表失败")
	}
	//查询用户文件列表总数
	count, err = l.svcCtx.Engine.Where("parent_id = ? AND user_identity = ?", req.Id, userIdentity).Count(new(models.UserRepository))
	if err != nil {
		return nil, errors.New("查询用户文件总数失败")
	}
	resp.List = userFile
	resp.Count = count

	return
}
