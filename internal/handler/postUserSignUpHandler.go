package handler

import (
	"net/http"

	"user/internal/logic"
	"user/internal/svc"
	"user/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func PostUserSignUpHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SignUpReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewPostUserSignUpLogic(r.Context(), svcCtx)
		resp, err := l.PostUserSignUp(r, w, req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
