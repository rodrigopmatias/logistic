package router

import "net/http"

type Context struct {
	Request *http.Request
}
