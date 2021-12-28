// Package dbagent package dbagent
package dbagent

import (
	"time"

	"gorm.io/gorm"
)

// HistoryTick HistoryTick
type HistoryTick struct {
	gorm.Model `json:"-" swaggerignore:"true"`

	Stock    *Stock    `json:"stock,omitempty" yaml:"stock" gorm:"foreignKey:StockID"`
	StockID  int64     `json:"stock_id,omitempty" yaml:"stock_id" gorm:"column:stock_id"`
	TickTime time.Time `json:"tick_time,omitempty" yaml:"tick_time" gorm:"column:tick_time"`

	Close     float64 `json:"close,omitempty" yaml:"close" gorm:"column:close"`
	TickType  int64   `json:"tick_type,omitempty" yaml:"tick_type" gorm:"column:tick_type"`
	Volume    int64   `json:"volume,omitempty" yaml:"volume" gorm:"column:volume"`
	BidPrice  float64 `json:"bid_price,omitempty" yaml:"bid_price" gorm:"column:bid_price"`
	BidVolume int64   `json:"bid_volume,omitempty" yaml:"bid_volume" gorm:"column:bid_volume"`
	AskPrice  float64 `json:"ask_price,omitempty" yaml:"ask_price" gorm:"column:ask_price"`
	AskVolume int64   `json:"ask_volume,omitempty" yaml:"ask_volume" gorm:"column:ask_volume"`
}

// InsertHistoryTick InsertHistoryTick
func (c *DBAgent) InsertHistoryTick(record *HistoryTick) error {
	err := c.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&record).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// InsertMultiHistoryTick InsertMultiHistoryTick
func (c *DBAgent) InsertMultiHistoryTick(recordArr []*HistoryTick) error {
	err := c.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.CreateInBatches(&recordArr, multiInsertBatchSize).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteAllHistoryTick DeleteAllHistoryTick
func (c *DBAgent) DeleteAllHistoryTick() error {
	err := c.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Not("id = 0").Unscoped().Delete(&HistoryTick{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// CheckHistoryTickExistByStockNum CheckHistoryTickExistByStockNum
func (c *DBAgent) CheckHistoryTickExistByStockNum(date time.Time) (bool, error) {
	var count int64
	if err := c.DB.Model(&HistoryTick{}).Where("tick_time >= ? and tick_time < ?", date, date.AddDate(0, 0, 1)).Count(&count).Error; err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

// HistoryTickArr HistoryTickArr
type HistoryTickArr []*HistoryTick

// GetTotalTime GetTotalTime
func (c HistoryTickArr) GetTotalTime() int64 {
	if len(c) > 1 {
		return c[len(c)-1].TickTime.Unix() - c[0].TickTime.Unix()
	}
	return 0
}

// GetHistoryTickByStockIDAndDate GetHistoryTickByStockIDAndDate
func (c *DBAgent) GetHistoryTickByStockIDAndDate(stockID int64, date time.Time) (HistoryTickArr, error) {
	var tmp HistoryTickArr
	err := c.DB.Preload("Stock").Model(&HistoryTick{}).Where("stock_id = ? and tick_time >= ? and tick_time < ?", stockID, date, date.AddDate(0, 0, 1)).Find(&tmp).Error
	return tmp, err
}
