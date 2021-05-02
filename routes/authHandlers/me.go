package authHandlers

import (
	"encoding/json"
	"net/http"

	"github.com/rodrigopmatias/ligistic/framework/router"
	"github.com/rodrigopmatias/ligistic/framework/router/context"
)

func MeHandler(ctx *context.Context) router.Result {
	user := ctx.Values["user"]
	var statusCode = http.StatusNotImplemented
	var content = []byte(`{"ok": false, "message": "not implmented"}`)

	if user != nil {
		statusCode = 200
		content, _ = json.Marshal(user)
	} else {
		statusCode = 404
		content = []byte(`{"ok": false, "message": "user not authenticated"}`)
	}

	return router.Result{
		StatusCode: statusCode,
		Content:    content,
	}
}
