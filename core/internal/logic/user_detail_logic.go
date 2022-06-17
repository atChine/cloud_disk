package logic

import (
	"cloud_disk/core/models"
	"context"
	"errors"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.DetailRequest) (resp *types.DetailReplay, err error) {
	// todo: add your logic here and delete this line
	//查询用户详情
	userDetail := new(models.UserBasic)
	get, err := l.svcCtx.Engine.Where("identity = ?", req.Identity).Get(userDetail)
	if err != nil {
		return nil, err
	}
	if !get {
		return nil,errors.New("no find user")
	}
	resp = new(types.DetailReplay)
	resp.Name = userDetail.Name
	resp.Email = userDetail.Email
	return
}
