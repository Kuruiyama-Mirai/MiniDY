package handler

import (
	"net/http"
	"strconv"

	"MiniDY/app/interaction/cmd/api/internal/logic"
	"MiniDY/app/interaction/cmd/api/internal/svc"
	"MiniDY/app/interaction/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func postCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentActionReq
		int_videoId, _ := strconv.Atoi(r.URL.Query().Get("video_id"))
		int_actionType, _ := strconv.Atoi(r.URL.Query().Get("action_type"))
		int_commentId, _ := strconv.Atoi(r.URL.Query().Get("comment_id"))

		req.VideoId = int64(int_videoId)
		req.ActionType = int64(int_actionType)
		req.CommentId = int64(int_commentId)

		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewPostCommentLogic(r.Context(), svcCtx)
		resp, err := l.PostComment(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
