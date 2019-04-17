package main

import (
	"github.com/lbrictson/stamp/internal/api"
	"github.com/lbrictson/stamp/internal/logging"
)

// main starts stamp
func main() {
	logging.InitLoggers()
	api.RunServer()
}
