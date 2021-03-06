package logic

import (
	"context"
	"math/rand"

	"user/internal/svc"
	"user/internal/types"
	"user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostUserSentSMSLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostUserSentSMSLogic(ctx context.Context, svcCtx *svc.ServiceContext) PostUserSentSMSLogic {
	return PostUserSentSMSLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostUserSentSMSLogic) PostUserSentSMS(req types.SentCheckCodeReq) (resp *types.ErrorRsb, err error) {
	// todo: add your logic here and delete this line

	// var errPhoneOccupied error = errors.New("phone")

	id := model.AccountToID(req.CountryCode + req.Phone)
	u, err := l.svcCtx.UsersModel.FindOne(id)
	//判断用户是否已经注册过了
	if u != nil && err == nil {
		return &types.ErrorRsb{Error: "phone"}, nil
	}
	//code:=120101
	mobile := req.CountryCode + req.Phone
	code := rand.Intn(900000) + 100000
	SMSCode[mobile] = code
	return &types.ErrorRsb{Error: "false"}, nil
	// err = SendSMS(mobile, strconv.Itoa(code))
	// if err == nil {
	// 	return &types.ErrorRsb{Error: "false"}, nil
	// }
	// return &types.ErrorRsb{Error: "true"}, nil
}

var SMSCode = make(map[string]int)

// SendSMS 发送短信
func SendSMS(mobile, Code string) (err error) {
	return nil
	// client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "你的AccessKey ID", "你的AccessKey Secret")
	// request := dysmsapi.CreateSendSmsRequest()
	// request.Scheme = "https"
	// request.PhoneNumbers = mobile
	// request.SignName = "短信开头"
	// request.TemplateCode = "短信模板代码"
	// request.TemplateParam = fmt.Sprintf("{code:%s}", Code)
	// response, err := client.SendSms(request)
	// if err != nil {
	// 	fmt.Print(err.Error())
	// }
	// fmt.Printf("ali sendmsg response code is %#v\n", response.BaseResponse.GetHttpStatus())
	// if response.BaseResponse.GetHttpStatus() == 200 {
	// 	return nil
	// }
	// return errors.New("fail")
}
func SendEmail(account string, Code string) (err error) {
	return nil
	// todo :fullfill code here
}
