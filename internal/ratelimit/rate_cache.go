package ratelimit

import (
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

var cacheDirectory map[string]*cache.Cache
var lock = sync.RWMutex{}

// init will spawn the intial empty cache map
func initCache() {
	m := make(map[string]*cache.Cache)
	cacheDirectory = m
	return
}

// createCache will create a new rate limiting cache with the expiration time in minutes
func createCache(minutes int) *cache.Cache {
	rateCache := cache.New(time.Duration(minutes)*time.Minute, time.Duration(minutes)*time.Minute)
	return rateCache
}

// incrementCache will increase the hit count by one for the IP in the cache and return true if it has hit the limit
func incrementCache(targetCache *cache.Cache, IPaddress string, limit int) bool {
	current := getRateCache(targetCache, IPaddress)
	if current == 0 {
		targetCache.Set(IPaddress, current+1, cache.DefaultExpiration)
		return false
	}
	targetCache.IncrementInt(IPaddress, 1)
	if current+1 > limit {
		return true
	}
	return false
}

// GetRateCache gets the rules for the specified host, the first return indicates
// if the hostname was in the cache
func getRateCache(targetCache *cache.Cache, IPAddress string) int {
	val, found := targetCache.Get(IPAddress)
	if found {
		return val.(int)
	}
	return 0
}
