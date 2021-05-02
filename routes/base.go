package routes

import (
	"github.com/rodrigopmatias/ligistic/framework/router"
	"github.com/rodrigopmatias/ligistic/routes/authHandlers"
	"github.com/rodrigopmatias/ligistic/routes/coreHandlers"
)

func Setup() {
	router.Register("GET", "/health", coreHandlers.HealthHandler)
	router.Register("GET", "/ping", coreHandlers.PingHandler)
	router.Register("POST", "/api/auth/register", authHandlers.RegisterHandle)
	router.Register("POST", "/api/auth/authenticate", authHandlers.AuthenticateHandle)
	router.Register("GET", "/api/auth/me", authHandlers.MeHandler)
}
