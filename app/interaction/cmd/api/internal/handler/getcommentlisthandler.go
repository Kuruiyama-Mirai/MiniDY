package handler

import (
	"net/http"

	"MiniDY/app/interaction/cmd/api/internal/logic"
	"MiniDY/app/interaction/cmd/api/internal/svc"
	"MiniDY/app/interaction/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func getCommentListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetCommentListLogic(r.Context(), svcCtx)
		resp, err := l.GetCommentList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
