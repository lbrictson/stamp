package main

import (
	"github.com/lbrictson/stamp/internal/api"
	"github.com/lbrictson/stamp/internal/logging"
	"github.com/lbrictson/stamp/internal/rulecache"
)

// main starts stamp
func main() {
	logging.Init()
	rulecache.Init()
	api.RunServer()
}
