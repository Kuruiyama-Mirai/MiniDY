package handler

import (
	"net/http"

	"MiniDY/app/videos/cmd/api/internal/logic"
	"MiniDY/app/videos/cmd/api/internal/svc"
	"MiniDY/app/videos/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func publistedListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishedVideoListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewPublistedListLogic(r.Context(), svcCtx)
		resp, err := l.PublistedList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
