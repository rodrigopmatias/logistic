package context

import "net/http"

type Context struct {
	Request *http.Request
	Values  map[string]interface{}
}

func New(req *http.Request) *Context {
	return &Context{
		Request: req,
		Values:  make(map[string]interface{}),
	}
}
