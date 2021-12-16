// Package cache package cache
package cache

import (
	"fmt"
	"trade_agent/pkg/dbagent"
)

// KeyStockDetail KeyStockDetail
func KeyStockDetail(stockNum string) string {
	return fmt.Sprintf("StockDetail:%s", stockNum)
}

// GetStock GetStock
func (c *Cache) GetStock(stockNum string) *dbagent.Stock {
	defer c.lock.RUnlock()
	c.lock.RLock()
	if value, ok := c.Cache.Get(KeyStockDetail(stockNum)); ok {
		return value.(*dbagent.Stock)
	}
	return nil
}