package chat

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"wx-chat/service/chat/cmd/api/internal/logic/chat"
	"wx-chat/service/chat/cmd/api/internal/svc"
	"wx-chat/service/chat/cmd/api/internal/types"
)

func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := chat.NewListLogic(r.Context(), svcCtx)
		resp, err := l.List(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
