package logic

import (
	"cloud_disk/core/define"
	"cloud_disk/core/helper"
	"context"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshAuthorizationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshAuthorizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshAuthorizationLogic {
	return &RefreshAuthorizationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshAuthorizationLogic) RefreshAuthorization(req *types.RefreshAuthorizationRequest, authorization string) (resp *types.RefreshAuthorizationReply, err error) {
	//解析token
	token, err := helper.AnalyzeToken(authorization)
	if err != nil {
		return nil, err
	}
	//生成新的token
	newToken, err := helper.GenerateToken(token.Id, token.Identity, token.Name, define.TokenExpire)
	if err != nil {
		return nil, err
	}
	//生成新的刷新token
	newRefreshToken, err := helper.GenerateToken(token.Id, token.Identity, token.Name, define.RefreshTokenExpire)
	if err != nil {
		return nil, err
	}
	resp = new(types.RefreshAuthorizationReply)
	resp.Token = newToken
	resp.RefreshToken = newRefreshToken
	return
}
