package handler

import (
	"net/http"

	"book/service/user/api/internal/logic"
	"book/service/user/api/internal/svc"
	"book/service/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func searchHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSearchLogic(r.Context(), svcCtx)
		resp, err := l.Search(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
