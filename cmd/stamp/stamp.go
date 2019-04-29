package main

import (
	"fmt"
	"os"

	"github.com/lbrictson/stamp/internal/api"
	"github.com/lbrictson/stamp/internal/logging"
	"github.com/lbrictson/stamp/internal/ratelimit"
	"github.com/lbrictson/stamp/internal/rulecache"
)

// main starts stamp
func main() {
	if len(os.Args) != 1 {
		fmt.Println("You must supply an argument to start")
	}
	arg := os.Args[1]
	switch arg {
	case "api":
		logging.Init()
		rulecache.Init()
		api.RunServer()
	case "rate":
		logging.Init()
		ratelimit.RunServer()
	}
}
