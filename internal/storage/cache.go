package storage

import "sync"

type OrderCache struct {
	m      sync.RWMutex
	orders map[string]Order
}

func NewCache() *OrderCache {
	return &OrderCache{
		orders: make(map[string]Order),
	}
}

func (c *OrderCache) GetOrd(OrderUID string) (Order, bool) {
	c.m.Lock()
	defer c.m.Unlock()
	order, ok := c.orders[OrderUID]
	return order, ok
}

func (c *OrderCache) SaveOrder(order Order) {
	c.m.RLock()
	defer c.m.RUnlock()
	c.orders[order.OrderUID] = order
}
