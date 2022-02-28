package logic

import (
	"context"

	"user/internal/svc"
	"user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostUserPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostUserPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) PostUserPasswordLogic {
	return PostUserPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostUserPasswordLogic) PostUserPassword(req types.ResetReq) (resp *types.ErrorRsb, err error) {
	// todo: add your logic here and delete this line

	return
}
