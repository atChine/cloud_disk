package logic

import (
	"cloud_disk/core/define"
	"cloud_disk/core/helper"
	"cloud_disk/core/models"
	"context"
	"errors"
	"time"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MailCodeSandRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCodeSandRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailCodeSandRegisterLogic {
	return &MailCodeSandRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailCodeSandRegisterLogic) MailCodeSandRegister(req *types.MailCodeSendRequest) (resp *types.MailCodeSendReplay, err error) {
	// todo: add your logic here and delete this line
	//只有该邮箱在数据中不存在
	count, err := l.svcCtx.Engine.Where("email = ?", req.Email).Count(new(models.UserBasic))
	if err != nil {
		return nil, err
	}
	if count > 0 {
		err = errors.New("该邮箱已被注册")
		return
	}
	//自动生成随机验证码
	code := helper.RandCode()
	//存储验证码到Redis
	l.svcCtx.RDB.Set(l.ctx,req.Email,code,time.Second * time.Duration(define.CodeExpire))
	//发送验证码
	err = helper.MailCodeSend(req.Email, code)
	if err != nil {
		return nil, err
	}
	return
}
