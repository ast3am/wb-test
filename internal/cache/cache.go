package cache

import (
	"github.com/example/internal/models"
)

type logger interface {
	DebugMsg(msg string)
}

type OrdersCache struct {
	cache map[string]models.Orders
	log   logger
}

func CacheInit(log logger) *OrdersCache {
	cacheMap := make(map[string]models.Orders)
	OrdersCache := OrdersCache{
		cache: cacheMap,
		log:   log,
	}
	log.DebugMsg("init cache is OK")
	return &OrdersCache
}

func (c *OrdersCache) InsertOrdersToCache(orders models.Orders) {
	c.cache[orders.Order.OrderUid] = orders
}

func (c *OrdersCache) GetOrdersById(Uid string) models.Orders {
	return c.cache[Uid]
}
