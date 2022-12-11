// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"go-zore-demo/user-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.TestMiddleWare},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/user-api/info",
					Handler: GetUserHandler(serverCtx),
				},
			}...,
		),
	)
}
