// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	CountryCode string `json:"countryCode"`
	Account     string `json:"account"`
	Password    string `json:"password"`
}

type SentCheckCodeReq struct {
	CountryCode string `json:"countryCode"`
	Phone       string `json:"phone"`
}

type SignUpReq struct {
	CountryCode string `json:"countryCode"`
	Account     string `json:"account"`
	Phone       string `json:"phone"`
	Password    string `json:"password"`
	SMSCode     int    `json:"SMSCode"`
}

type ResetReq struct {
	CountryCode string `json:"countryCode"`
	Account     string `json:"account"`
	Password    string `json:"password"`
	CheckCode   int    `json:"checkCode"`
}

type ErrorRsb struct {
	Error string `json:"error"`
}

type JwtReq struct {
	Width       uint16 `form:"width"`
	Height      uint16 `form:"height"`
	AvailWidth  uint16 `form:"availWidth"`
	AvailHeight uint16 `form:"availHeight"`
	OuterHeight uint16 `form:"outerHeight"`
	OuterWidth  uint16 `form:"outerWidth"`
	InnerHeight uint16 `form:"innerHeight"`
	InnerWidth  uint16 `form:"innerWidth"`
}

type JwtRsp struct {
	Jwt              string `json:"jwt"`
	Id               string `json:"id"`
	Sub              string `json:"sub"`
	TemporaryAccount bool   `json:"temporaryAccount"`
}

type AccountOccupiedReq struct {
	Name string `form:"name"`
}

type AccountOccupiedRsb struct {
	Error string `json:"error"`
}

type AccountID struct {
	Id string `form:"id"`
}

type NameRsp struct {
	Name string `json:"name"`
}

type JwtReturn struct {
	Jwt string `json:"jwt"`
}

type MyProfileReturn struct {
	Succ         bool   `json:"succ"`
	CountryCode  string `json:"countryCode"`
	Phone        string `json:"phone"`
	Nick         string `json:"nickName"`
	Introduction string `json:"introduction"`
	Avatar       string `json:"avatar"`
	RealName     string `json:"realName"`
}

type MyProfileChangeReq struct {
	ChannelName  string `form:"channelName"`
	LoginAccount string `form:"loginAccount"`
}
