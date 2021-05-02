package router

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"
)

type ContentResult struct {
	Ok bool `json:"ok"`
}

type MessageContentResult struct {
	ContentResult
	Message string `json:"message"`
}

type Result struct {
	StatusCode int
	Content    []byte
}

type Route struct {
	Method  string
	Pattern *regexp.Regexp
	Handle  func(ctx *Context) Result
}

var routes []Route

func Register(method string, pattern string, handle func(ctx *Context) Result) error {
	compiled, err := regexp.Compile(pattern)

	if err == nil {
		routes = append(routes, Route{
			Method:  method,
			Pattern: compiled,
			Handle:  handle,
		})
	} else {
		return err
	}

	return nil
}

func RouterHandler(rw http.ResponseWriter, req *http.Request) {
	statusCode := 404
	content := []byte("{\"ok\": false, \"message\": \"Resource not found\"}")

	ctx := Context{
		Request: req,
	}

	start := time.Now()
	for _, route := range routes {
		if route.Pattern.MatchString(req.URL.Path) {
			result := route.Handle(&ctx)
			statusCode = result.StatusCode
			content = result.Content
			break
		}
	}
	end := time.Now()

	elapsedStr := elapsed(int64(end.Sub(start)))
	log.Printf("%s - %s - %d - %s", req.Method, req.URL.Path, statusCode, elapsedStr)

	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("Elapsed", elapsedStr)
	rw.WriteHeader(statusCode)
	rw.Write(content)
}

func elapsed(nanotime int64) string {
	units := []string{
		"ns",
		"us",
		"ms",
		"s",
	}

	unit := "ns"
	current := float64(nanotime)

	for _, currentUnit := range units {
		if current > 1000.0 {
			current = current / 1000.0
		} else {
			unit = currentUnit
			break
		}
	}

	return fmt.Sprintf("%0.1f %s", current, unit)
}
