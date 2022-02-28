package logic

import (
	"context"
	"net/http"

	"user/internal/svc"
	"user/internal/types"

	"github.com/yangkequn/GoTools"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserLogic {
	return GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(w http.ResponseWriter, req types.AccountID) (resp *types.NameRsp, err error) {
	id := GoTools.StringToInt64(req.Id)
	if id == 0 {
		return &types.NameRsp{Name: ""}, nil
	}
	user, err := l.svcCtx.UsersModel.FindOne(id)
	w.Header().Add("Cache-Control", "public,max-age=86400")
	if err == nil && user.Id != 0 {
		return &types.NameRsp{Name: user.Nick}, nil
	}
	return &types.NameRsp{Name: ""}, nil
}
