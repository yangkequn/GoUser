package handler

import (
	"net/http"

	"user/internal/logic"
	"user/internal/svc"
	"user/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserAccountOccupiedHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AccountOccupiedReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetUserAccountOccupiedLogic(r.Context(), svcCtx)
		resp, err := l.GetUserAccountOccupied(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
