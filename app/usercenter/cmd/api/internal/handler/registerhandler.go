package handler

import (
	"net/http"

	"MiniDY/app/usercenter/cmd/api/internal/logic"
	"MiniDY/app/usercenter/cmd/api/internal/svc"
	"MiniDY/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func registerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq

		req.Username = r.URL.Query().Get("username")
		req.Password = r.URL.Query().Get("password")

		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
