package api

import (
	"github.com/lbrictson/stamp/internal/logging"
	"github.com/lbrictson/stamp/internal/rules"

	"github.com/labstack/echo"
)

// RunServer will start the api server
func RunServer() {
	e := echo.New()
	e.Any("/authz", authzRoute)
	e.GET("/heartbeat", heartbeat)
	e.Logger.Fatal(e.Start(":8667"))
}

func authzRoute(c echo.Context) error {
	return c.String(200, "ok")
}

// heartbeat always returns a 200 status code and ok, used for verifying the
// api server is running
func heartbeat(c echo.Context) error {
	logging.Logger.Info("Heartbeat hit")
	// Remove later, just for demoing
	dummyCheck := rules.HeaderRule{
		Name:          "silly-heartbeat-rule",
		Header:        "Sample-H",
		Value:         "test",
		ExactMatch:    true,
		WhiteListed:   true,
		CaseSensitive: true,
	}
	block := dummyCheck.Eval(c.Request())
	if block {
		return c.String(501, "Unauthorized")
	}
	return c.String(200, "ok")
}
