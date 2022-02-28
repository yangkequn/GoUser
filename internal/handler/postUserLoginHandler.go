package handler

import (
	"net/http"

	"user/internal/logic"
	"user/internal/svc"
	"user/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func PostUserLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewPostUserLoginLogic(r.Context(), svcCtx)
		resp, err := l.PostUserLogin(r, w, req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
