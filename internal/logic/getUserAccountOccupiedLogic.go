package logic

import (
	"context"

	"user/internal/svc"
	"user/internal/types"
	"user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAccountOccupiedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserAccountOccupiedLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserAccountOccupiedLogic {
	return GetUserAccountOccupiedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserAccountOccupiedLogic) GetUserAccountOccupied(req types.AccountOccupiedReq) (resp *types.AccountOccupiedRsb, err error) {
	id := model.AccountToID(req.Name)
	if id == 0 {
		return &types.AccountOccupiedRsb{Error: "账号无效"}, nil
	}
	if u, err := l.svcCtx.UsersModel.FindOne(id); err == nil && u.Id != 0 {
		return &types.AccountOccupiedRsb{Error: "账号已经存在"}, nil
	}
	return &types.AccountOccupiedRsb{Error: ""}, nil
}
