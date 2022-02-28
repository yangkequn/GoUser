package logic

import (
	"context"
	"net/http"
	"time"

	"user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserSignOutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserSignOutLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserSignOutLogic {
	return GetUserSignOutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserSignOutLogic) GetUserSignOut(r *http.Request, w http.ResponseWriter) error {
	cookie := http.Cookie{Name: "Authorization", Value: "deleted", Path: "/", Expires: time.Now().Add(-time.Hour * 24), HttpOnly: true}
	http.SetCookie(w, &cookie)
	return nil

}
