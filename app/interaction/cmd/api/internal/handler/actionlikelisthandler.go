package handler

import (
	"net/http"
	"strconv"

	"MiniDY/app/interaction/cmd/api/internal/logic"
	"MiniDY/app/interaction/cmd/api/internal/svc"
	"MiniDY/app/interaction/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func actionLikeListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InteractionListReq
		int_userId, _ := strconv.Atoi(r.URL.Query().Get("user_id"))

		req.UserId = int64(int_userId)
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewActionLikeListLogic(r.Context(), svcCtx)
		resp, err := l.ActionLikeList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
