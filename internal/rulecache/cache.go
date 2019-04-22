package rulecache

import (
	"github.com/lbrictson/stamp/internal/rules"
	"github.com/patrickmn/go-cache"
)

// ruleCache is a map of hostnames and their rules, refreshed every 60 seconds
var ruleCache *cache.Cache

// Init creates the hostname rule cache and spawns the
// kubernetes goroutine that manages keeping it fresh
func Init() {
	ruleCache = cache.New(cache.NoExpiration, cache.NoExpiration)
}

// UpdateCacheHostRules updates OR creates a host value in the cache with the specified rules
func UpdateCacheHostRules(hostname string, ruleList []rules.Rule) error {
	ruleCache.Set(hostname, ruleList, cache.NoExpiration)
	return nil
}

// GetCacheHostRules gets the rules for the specified host, the first return indicates
// if the hostname was in the cache
func GetCacheHostRules(hostname string) (bool, []rules.Rule, error) {
	ruleList, found := ruleCache.Get(hostname)
	if found {
		return found, ruleList.([]rules.Rule), nil
	}
	return found, []rules.Rule{}, nil

}

// DeleteCacheHostRules deletes the hostname and rules associated with it from the cache
func DeleteCacheHostRules(hostname string) error {
	ruleCache.Delete(hostname)
	return nil
}
