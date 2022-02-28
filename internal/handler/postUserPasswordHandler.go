package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"user/internal/logic"
	"user/internal/svc"
	"user/internal/types"
)

func PostUserPasswordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ResetReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewPostUserPasswordLogic(r.Context(), svcCtx)
		resp, err := l.PostUserPassword(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
