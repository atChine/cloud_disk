package logic

import (
	"context"
	"errors"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"

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
	return
}
