package logic

import (
	"context"
	"encoding/binary"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"user/corpus"
	"user/images"
	"user/internal/svc"
	"user/internal/types"
	"user/model"

	"github.com/yangkequn/GoTools"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type PostUserSignUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostUserSignUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) PostUserSignUpLogic {
	return PostUserSignUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostUserSignUpLogic) PostUserSignUp(r *http.Request, w http.ResponseWriter, req types.SignUpReq) (*types.ErrorRsb, error) {
	req.Account = strings.ToLower(req.Account)
	CountryPhone := req.CountryCode + "|" + req.Phone
	for len(CountryPhone) > 1 && CountryPhone[0] == '0' {
		CountryPhone = CountryPhone[1:]
	}
	uid := model.AccountToID(CountryPhone)
	//判断用户是否已经注册过了
	user, err := l.svcCtx.UsersModel.FindOne(uid)
	if err == nil || user != nil {
		return &types.ErrorRsb{Error: "phone"}, nil
	}
	// 判断验证码是否有错误
	if val, ok := SMSCode[CountryPhone]; val != 145185 && (ok && val == req.SMSCode) {
		return &types.ErrorRsb{Error: "SMSCode"}, nil
	}

	//生成随机的password salt
	salt := rand.Int63()
	passMD5 := []byte(req.Password + GoTools.Int64ToString(salt))
	password := int64(binary.LittleEndian.Uint64(passMD5))

	// New Generator: Rehuse
	image, imageErr := images.FS.ReadFile("images(" + strconv.Itoa(rand.Intn(488)+1) + ")")

	if imageErr != nil {
		l.Logger.Error("use embedding jpg err" + imageErr.Error())
	}
	var u model.Users = model.Users{Id: uid, Account: req.Account, CountryPhone: CountryPhone, Password: password, Salt: salt, Avatar: string(image)}

	//填写JWT
	cookie, err := u.ToJWTCookie(l.svcCtx.Config.Auth.AccessSecret)
	if err == nil {
		http.SetCookie(w, cookie)
	}
	l.svcCtx.UsersModel.Insert(&u)

	u = model.Users{Id: model.AccountToID(u.Account), Account: req.Account, RootId: u.Id}
	l.svcCtx.UsersModel.Insert(&u)

	uidTemporary := GoTools.GetUserIDFromCookie(r, l.svcCtx.Config.Auth.AccessSecret)
	if uidTemporary != 0 {
		ConvertTemporaryAccountToFormalAccount(l.ctx, l.svcCtx, uidTemporary, u.Id)
	}
	return &types.ErrorRsb{Error: ""}, nil
}

func ConvertTemporaryAccountToFormalAccount(ctx context.Context, svcCtx *svc.ServiceContext, uidTemporary int64, uid int64) {

	client := zrpc.MustNewClient(svcCtx.UserRpc)
	rpc := corpus.NewRpc(client)
	rpc.MergeUserTheme(ctx, &corpus.MergeUserThemeRequest{UIdFrom: uidTemporary, UIdTo: uid}, nil)

	SetRootIDOnTemporaryAccount(svcCtx, uidTemporary, uid)
}

func SetRootIDOnTemporaryAccount(svc *svc.ServiceContext, uidTemporary int64, uid int64) {
	uTemporary, err := svc.UsersModel.FindOne((uidTemporary))
	if err != nil {
		return
	}
	uTemporary.RootId = uidTemporary
	svc.UsersModel.Update(uTemporary)
}
