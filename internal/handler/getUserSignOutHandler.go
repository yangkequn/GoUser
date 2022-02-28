package handler

import (
	"net/http"

	"user/internal/logic"
	"user/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserSignOutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetUserSignOutLogic(r.Context(), svcCtx)
		err := l.GetUserSignOut(r, w)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
