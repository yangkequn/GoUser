package logic

import (
	"context"
	"strings"

	"user/internal/svc"
	"user/internal/types"
	"user/model"

	"github.com/yangkequn/GoTools"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserProfileLogic {
	return GetUserProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func UID(ctx context.Context) (int64, error) {
	UId := ctx.Value("id").(string)
	if len(UId) == 0 {
		return int64(0), model.ErrLoginNeeded
	}
	return GoTools.StringToInt64(UId), nil
}
func (l *GetUserProfileLogic) GetUserProfile() (resp *types.MyProfileReturn, err error) {
	uid, uErr := UID(l.ctx)
	if uErr != nil {
		return nil, uErr
	}
	u, err := l.svcCtx.UsersModel.FindOne(uid)
	if u.Id == 0 || err != nil {
		return &types.MyProfileReturn{Succ: false}, nil
	}
	countryPhone := strings.Split(u.CountryPhone, "|")
	country, phone := countryPhone[0], countryPhone[1]
	return &types.MyProfileReturn{Succ: true, CountryCode: country, Phone: phone, Introduction: u.Introduction, RealName: ""}, nil
}
