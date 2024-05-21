package utils

import (
	"time"

	"github.com/patrickmn/go-cache"
)

// Cache is a global variable that holds the cache instance
var Cache *cache.Cache

func InitCache() {
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	Cache = cache.New(5*time.Minute, 10*time.Minute)
}
