package logic

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"user/internal/svc"
	"user/internal/types"
	"user/model"

	"github.com/yangkequn/GoTools"
	"github.com/zeromicro/go-zero/core/logx"
)

type PostUserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) PostUserLoginLogic {
	return PostUserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostUserLoginLogic) PostUserLogin(r *http.Request, w http.ResponseWriter, req types.LoginReq) (*types.ErrorRsb, error) {
	req.Account = strings.ToLower(req.Account)
	id := model.AccountToID(req.CountryCode + "|" + req.Account)
	u, err := l.svcCtx.UsersModel.FindOne(id)
	if err != nil || u == nil {
		u, err = l.svcCtx.UsersModel.FindOne(model.AccountToID(req.Account))
		if err != nil || u == nil {
			return &types.ErrorRsb{Error: ""}, model.ErrLoginFail
		}
	}
	//一个账号可以有多个账号名，这些账号最终需要追溯到原始根账号
	if u.RootId != 0 {
		u, err = l.svcCtx.UsersModel.FindOne(u.RootId)
		if err != nil || u == nil {
			return &types.ErrorRsb{Error: ""}, errors.New("LoginFail")
		}
	}

	//把当前临时账号的内容保存到被登录账号
	uidTemporary := GoTools.GetUserIDFromCookie(r, l.svcCtx.Config.Auth.AccessSecret)
	if uidTemporary != 0 {
		ConvertTemporaryAccountToFormalAccount(l.ctx, l.svcCtx, uidTemporary, u.Id)
	}

	cookie, err := u.ToJWTCookie(l.svcCtx.Config.Auth.AccessSecret)
	if err == nil {
		http.SetCookie(w, cookie)
	}
	return &types.ErrorRsb{Error: ""}, nil
}
