package handler

import (
	"net/http"

	"MiniDY/app/social/cmd/api/internal/logic"
	"MiniDY/app/social/cmd/api/internal/svc"
	"MiniDY/app/social/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func followerListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FollowerListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFollowerListLogic(r.Context(), svcCtx)
		resp, err := l.FollowerList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
