package authHandlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/rodrigopmatias/ligistic/controllers/authController"
	"github.com/rodrigopmatias/ligistic/framework/router"
)

func AuthenticateHandle(ctx *router.Context) router.Result {
	var statusCode = http.StatusNotImplemented
	var content = []byte("{\"ok\": false, \"message\": \"not implemented\"}")
	var payload = authController.AuthenticatePayload{}

	body, err := ioutil.ReadAll(ctx.Request.Body)

	if err != nil {
		statusCode = 400
		content = []byte("{\"ok\": false, \"message\": \"can't ready request body\"}")
	} else {
		err = json.Unmarshal(body, &payload)

		if err != nil {
			statusCode = 400
			content = []byte("{\"ok\": false, \"message\": \"invlaid payload\"}")
		} else {
			result, err := authController.Authenticate(&payload)

			if err != nil {
				statusCode = http.StatusBadRequest
				content, _ = json.Marshal(router.MessageContentResult{
					ContentResult: router.ContentResult{
						Ok: true,
					},
					Message: err.Error(),
				})
			} else {
				statusCode = http.StatusCreated
				content, _ = json.Marshal(result)
			}

		}
	}

	return router.Result{
		StatusCode: statusCode,
		Content:    content,
	}
}
