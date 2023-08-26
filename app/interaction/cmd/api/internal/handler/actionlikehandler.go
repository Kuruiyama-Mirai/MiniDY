package handler

import (
	"net/http"
	"strconv"

	"MiniDY/app/interaction/cmd/api/internal/logic"
	"MiniDY/app/interaction/cmd/api/internal/svc"
	"MiniDY/app/interaction/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func actionLikeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InteractionReq
		//做一个参数提取
		int_actionType, _ := strconv.Atoi(r.URL.Query().Get("action_type"))
		int_videoId, _ := strconv.Atoi(r.URL.Query().Get("video_id"))

		req.ActionType = int64(int_actionType)
		req.VideoId = int64(int_videoId)
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewActionLikeLogic(r.Context(), svcCtx)
		resp, err := l.ActionLike(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
