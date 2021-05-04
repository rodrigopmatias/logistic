package middleware

import (
	"github.com/rodrigopmatias/ligistic/framework/router/context"
)

type middlewareHandle func(ctx *context.Context)

type middlewareStack struct {
	beforer []middlewareHandle
	after   []middlewareHandle
}

var stack middlewareStack

func OnBefore(handle middlewareHandle) {
	stack.beforer = append(stack.beforer, handle)
}

func OnAfter(handle middlewareHandle) {
	stack.after = append(stack.after, handle)
}

func DoBefore(ctx *context.Context) {
	for _, handle := range stack.beforer {
		handle(ctx)
	}
}

func DoAfter(ctx *context.Context) {
	for _, handle := range stack.beforer {
		handle(ctx)
	}
}
