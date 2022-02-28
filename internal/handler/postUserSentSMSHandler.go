package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"user/internal/logic"
	"user/internal/svc"
	"user/internal/types"
)

func PostUserSentSMSHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SentCheckCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewPostUserSentSMSLogic(r.Context(), svcCtx)
		resp, err := l.PostUserSentSMS(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
