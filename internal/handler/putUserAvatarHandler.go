package handler

import (
	"net/http"

	"user/internal/logic"
	"user/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func putUserAvatarHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewPutUserAvatarLogic(r.Context(), svcCtx)
		err := l.PutUserAvatar(r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
