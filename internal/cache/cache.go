package cache

import "github.com/lbrictson/stamp/internal/rules"

// Cache is a map of hostnames and their rules, refreshed every 60 seconds
var Cache map[string][]rules.Rule

// InitCache creates the hostname rule cache and spawns the
// kubernetes goroutine that manages keeping it fresh
func InitCache() {
	Cache = make(map[string][]rules.Rule)
}
