package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"wx-chat/service/user/cmd/api/internal/logic/user"
	"wx-chat/service/user/cmd/api/internal/svc"
	"wx-chat/service/user/cmd/api/internal/types"
)

func FindUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FindUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewFindUserLogic(r.Context(), svcCtx)
		resp, err := l.FindUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
