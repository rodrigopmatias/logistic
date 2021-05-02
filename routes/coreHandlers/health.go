package coreHandlers

import (
	"encoding/json"

	"github.com/rodrigopmatias/ligistic/framework/router"
)

type HealthResult struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func HealthHandler(ctx *router.Context) router.Result {
	result := HealthResult{
		Ok:      true,
		Message: "Service is alive!",
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
