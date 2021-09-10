package router

import (
	"fmt"
	"os"

	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func New() *echo.Echo {
	e := echo.New()
	env := os.Getenv("ENVIRONMENT")

	if env == "dev" || env == "prod" {
		app, err := newrelic.NewApplication(
			newrelic.ConfigAppName("app_name"),
			newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
			newrelic.ConfigDebugLogger(os.Stdout),
		)
		if nil != err {
			fmt.Println("New Relic not integrated")
			fmt.Println(err)

		} else {

			// Newrelic needs to be the first Middleware that is added
			e.Use(nrecho.Middleware(app))

		}
		// Sentry Middleware
		e.Use(sentryecho.New(sentryecho.Options{Repanic: true}))

	}

	e.Logger.SetLevel(log.DEBUG)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, "sentry-trace"},
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowCredentials: true,
	}))

	return e
}
