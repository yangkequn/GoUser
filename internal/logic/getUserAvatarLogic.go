package logic

import (
	"context"
	"net/http"
	"user/internal/svc"
	"user/internal/types"

	"github.com/yangkequn/GoTools"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserAvatarLogic {
	return GetUserAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserAvatarLogic) GetUserAvatar(w http.ResponseWriter, req types.AccountID) error {
	uid := GoTools.StringToInt64(req.Id)
	if uid == 0 {
		return nil
	}

	u, er := l.svcCtx.UsersModel.FindOne(uid)
	if er != nil {
		return er
	}
	if len(u.Avatar) == 0 {
		return nil
	}
	avatar := []byte(u.Avatar)
	w.Header().Set("Content-Type", "image/webp")
	w.Write(avatar)
	return nil
}
