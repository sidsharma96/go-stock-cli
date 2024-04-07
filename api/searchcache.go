package api

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

type searchCallCache struct {
	stockData *cache.Cache
}

const (
	defaultExpiration = 24 * time.Hour
	purgeTime = 10 * time.Minute
)

func newCache() *searchCallCache {
    cache := cache.New(defaultExpiration, purgeTime)
    return &searchCallCache{
        stockData: cache,
    }
}

func (c *searchCallCache) read(id string) (TimeSeriesResponse, bool) {
    data, ok := c.stockData.Get(id)
    if ok {
        fmt.Println("from cache")
        res := data.(TimeSeriesResponse)
        return res, true
    }

    return TimeSeriesResponse{}, false
}

func (c *searchCallCache) update(id string, data TimeSeriesResponse) {
    c.stockData.Set(id, data, cache.DefaultExpiration)
}

var searchCache = newCache()