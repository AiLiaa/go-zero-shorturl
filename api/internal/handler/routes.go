// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/AiLiaa/go-zero-shorturl/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/shorten",
				Handler: ShortenHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/expand",
				Handler: ExpandHandler(serverCtx),
			},
		},
	)
}
