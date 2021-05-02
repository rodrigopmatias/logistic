package middleware

import (
	"log"

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
	log.Println("Do Before Handle")
	for _, handle := range stack.beforer {
		handle(ctx)
	}
}

func DoAfter(ctx *context.Context) {
	log.Println("Do After Handle")
	for _, handle := range stack.beforer {
		handle(ctx)
	}
}
