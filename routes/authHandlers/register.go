package authHandlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/rodrigopmatias/ligistic/controllers/authController"
	"github.com/rodrigopmatias/ligistic/framework/router"
)

func RegisterHandle(ctx *router.Context) router.Result {
	var statusCode = http.StatusNotImplemented
	var content []byte
	var payload = authController.RegisterPayload{}

	body, err := ioutil.ReadAll(ctx.Request.Body)

	if err != nil {
		statusCode = http.StatusBadRequest
		content = []byte("{\"ok\": false, \"message\": \"can't ready request body\"}")
	} else {
		err = json.Unmarshal(body, &payload)

		if err != nil {
			statusCode = http.StatusBadRequest
			content = []byte("{\"ok\": false, \"message\": \"invlaid payload\"}")
		} else {
			registerResult, err := authController.Register(&payload)

			if err != nil {
				statusCode = http.StatusBadRequest
				content, _ = json.Marshal(router.MessageContentResult{
					ContentResult: router.ContentResult{
						Ok: false,
					},
					Message: err.Error(),
				})
			} else {
				statusCode = http.StatusCreated
				content, _ = json.Marshal(registerResult)
			}
		}
	}

	return router.Result{
		StatusCode: statusCode,
		Content:    content,
	}
}
