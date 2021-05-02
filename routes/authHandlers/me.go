package authHandlers

import (
	"net/http"

	"github.com/rodrigopmatias/ligistic/framework/router"
)

func MeHandler(ctx *router.Context) router.Result {
	return router.Result{
		StatusCode: http.StatusNotImplemented,
		Content:    []byte("{\"ok\": false, \"message\": \"not implemented\"}"),
	}
}
