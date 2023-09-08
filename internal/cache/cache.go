package cache

import (
	"github.com/example/internal/models"
)

type OrdersCache struct {
	cache map[string]models.Orders
}

func CacheInit() *OrdersCache {
	cacheMap := make(map[string]models.Orders)
	OrdersCache := OrdersCache{
		cache: cacheMap,
	}
	return &OrdersCache
}

func (c *OrdersCache) InsertOrdersToCache(orders models.Orders) {
	c.cache[orders.Order.OrderUid] = orders
}

func (c *OrdersCache) GetOrdersById(Uid string) models.Orders {
	return c.cache[Uid]
}
