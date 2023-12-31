// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"MiniDY/app/videos/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/feed",
				Handler: videofeedHandler(serverCtx),
			},
		},
		rest.WithPrefix("/douyin"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/action",
				Handler: publishHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/list",
				Handler: publistedListHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/douyin/publish"),
	)
}
