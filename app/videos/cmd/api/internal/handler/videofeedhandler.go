package handler

import (
	"net/http"

	"MiniDY/app/videos/cmd/api/internal/logic"
	"MiniDY/app/videos/cmd/api/internal/svc"
	"MiniDY/app/videos/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func videofeedHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VideGetFeedReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewVideofeedLogic(r.Context(), svcCtx)
		resp, err := l.Videofeed(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
