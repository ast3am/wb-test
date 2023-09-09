package cache

import (
	"github.com/example/internal/models"
	"github.com/example/pkg/logging"
)

type OrdersCache struct {
	cache map[string]models.Orders
	log   *logging.Logger
}

func CacheInit(log *logging.Logger) *OrdersCache {
	cacheMap := make(map[string]models.Orders)
	OrdersCache := OrdersCache{
		cache: cacheMap,
		log:   log,
	}
	return &OrdersCache
}

func (c *OrdersCache) InsertOrdersToCache(orders models.Orders) {
	c.cache[orders.Order.OrderUid] = orders
}

func (c *OrdersCache) GetOrdersById(Uid string) models.Orders {
	return c.cache[Uid]
}
