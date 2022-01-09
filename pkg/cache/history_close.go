// Package cache package cache
package cache

import (
	"fmt"
	"time"
	"trade_agent/global"
)

// KeyStockHistoryClose KeyStockHistoryClose
func KeyStockHistoryClose(stockNum string, date time.Time) *Key {
	return &Key{
		Name: fmt.Sprintf("KeyStockHistoryClose:%s:%s", stockNum, date.Format(global.ShortTimeLayout)),
		Type: historyClose,
	}
}

// SetStockHistoryClose SetStockHistoryClose
func (c *Cache) SetStockHistoryClose(stockNum string, close float64, date time.Time) {
	key := KeyStockHistoryClose(stockNum, date)
	c.getCacheByType(key.Type).Set(key.Name, close, noExpired)
}

// GetHistoryClose GetHistoryClose
func (c *Cache) GetHistoryClose(stockNum string, date time.Time) float64 {
	c.lock.RLock()
	k := KeyStockHistoryClose(stockNum, date)
	tmp := c.CacheMap[string(k.Type)]
	c.lock.RUnlock()
	if tmp == nil {
		return 0
	}
	if value, ok := tmp.Get(k.Name); ok {
		return value.(float64)
	}
	return 0
}
