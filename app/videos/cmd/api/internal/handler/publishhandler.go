package handler

import (
	"net/http"

	"MiniDY/app/videos/cmd/api/internal/logic"
	"MiniDY/app/videos/cmd/api/internal/svc"
	"MiniDY/app/videos/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func publishHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishVideoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			return
		}

		//往logic里传递req
		l := logic.NewPublishLogic(r.Context(), svcCtx)
		resp, err := l.Publish(&req, file, fileHeader.Filename)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
