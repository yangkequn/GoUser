package logic

import (
	"context"

	"user/internal/svc"
	"user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutUserProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) PutUserProfileLogic {
	return PutUserProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutUserProfileLogic) PutUserProfile(req types.MyProfileChangeReq) (resp *types.ErrorRsb, err error) {
	uid, uErr := UID(l.ctx)
	if uErr != nil {
		return nil, uErr
	}
	u, err := l.svcCtx.UsersModel.FindOne(uid)
	if err != nil && u.Id != 0 {
		u.Nick = req.ChannelName
		u.Account = req.LoginAccount
		l.svcCtx.UsersModel.Update(u)
		return &types.ErrorRsb{Error: ""}, nil
	}
	return &types.ErrorRsb{Error: "user invalid"}, nil
}
