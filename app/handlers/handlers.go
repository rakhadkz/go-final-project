package handlers

import (
	"log"
	"mizu2281/mizus-goservice/business/auth"
	"mizu2281/mizus-goservice/business/mid"
	"mizu2281/mizus-goservice/foundation/web"
	"net/http"
	"os"
)

// API constructs an http.Handler with all application routes defined.
func API(build string, shutdown chan os.Signal, log *log.Logger, a *auth.Auth) *web.App {

	app := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log), mid.Metrics(), mid.Panics(log))

	check := check{
		log: log,
	}

	app.Handle(http.MethodGet, "/readiness", check.readiness, mid.Authenticate(a), mid.Authorize(auth.RoleAdmin))

	return app
}
