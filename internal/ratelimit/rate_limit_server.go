package ratelimit

import (
	"github.com/labstack/echo"
	"github.com/lbrictson/stamp/internal/logging"
)

// RunServer will start the api server
func RunServer() {
	initCache()
	e := echo.New()
	e.GET("/heartbeat", heartbeat)
	e.POST("/api/v1/limitip", limitIngressIP)
	e.Logger.Fatal(e.Start(":8668"))

}

// heartbeat always returns a 200 status code and ok, used for verifying the
// ratelimit server is running
func heartbeat(c echo.Context) error {
	logging.Logger.Info("Heartbeat hit")
	return c.String(200, "ok")
}

// limitIngressIP accepts a domain and an IP address
func limitIngressIP(c echo.Context) error {
	payload := new(IPPayload)
	if err := c.Bind(payload); err != nil {
		logging.Logger.Errorf("Error decoding payload %v", err)
		return c.String(500, err.Error())
	}
	if val, ok := cacheDirectory[payload.Domain]; ok {
		blockMe := incrementCache(val, payload.IPAddr, payload.Limit)
		if blockMe {
			logging.Logger.Infof("Rate limiting %v because breached limit of %v on domain %v", payload.IPAddr, payload.Limit, payload.Domain)
			return c.String(492, "Rate limited")
		}
	} else {
		lock.Lock()
		defer lock.Unlock()
		cacheDirectory[payload.Domain] = createCache(1)
		incrementCache(cacheDirectory[payload.Domain], payload.IPAddr, payload.Limit)
	}
	return c.String(200, "ok")
}
