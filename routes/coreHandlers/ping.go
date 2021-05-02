package coreHandlers

import (
	"encoding/json"
	"net/http"

	"github.com/rodrigopmatias/ligistic/framework/router"
)

type PingResult struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func PingHandler(req *http.Request) router.Result {
	result := PingResult{
		Ok:      true,
		Message: "Pong!",
	}

	content, err := json.Marshal(result)

	if err == nil {
		return router.Result{
			StatusCode: 200,
			Content:    content,
		}
	} else {
		return router.Result{
			StatusCode: 502,
			Content:    []byte("{\"ok\": false, \"message\", \"bad implementation\"}"),
		}
	}
}
