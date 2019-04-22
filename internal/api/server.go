package api

import (
	"github.com/lbrictson/stamp/internal/rulecache"

	"github.com/lbrictson/stamp/internal/logging"

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
	found, ruleList, _ := rulecache.GetCacheHostRules(c.Request().Referer())
	logging.Logger.Info("Auth hit")
	if found != true {
		logging.Logger.Infof("%v was not found in cache", c.Request().Referer())
		return c.String(200, "ok")
	}
	for _, x := range ruleList {
		logging.Logger.Infof("Checking rule %v", x.GetName())
		logging.Logger.Info(c.Request().RemoteAddr)
		block := x.Eval(c.Request())
		if block {
			logging.Logger.Info("Blocking request from %v to host %v because rule %v returned block", c.Request().RemoteAddr, c.Request().Referer(), x.GetName())
			return c.String(501, "Unauthorized")
		}
	}
	return c.String(200, "ok")
}

// heartbeat always returns a 200 status code and ok, used for verifying the
// api server is running
func heartbeat(c echo.Context) error {
	logging.Logger.Info("Heartbeat hit")
	return c.String(200, "ok")
}
