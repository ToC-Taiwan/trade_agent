// Package dbagent package dbagent
package dbagent

import (
	"time"

	"gorm.io/gorm"
)

// Target Target
type Target struct {
	gorm.Model `json:"-" swaggerignore:"true"`

	Stock    *Stock    `json:"stock" yaml:"stock" gorm:"foreignKey:StockID"`
	StockID  int64     `json:"stock_id" yaml:"stock_id" gorm:"column:stock_id"`
	TradeDay time.Time `json:"trade_day" yaml:"trade_day" gorm:"column:trade_day"`

	Rank        int   `json:"rank" yaml:"rank" gorm:"column:rank"`
	Volume      int64 `json:"volume" yaml:"volume" gorm:"column:volume"`
	Subscribe   bool  `json:"subscribe" yaml:"subscribe"`
	RealTimeAdd bool  `json:"real_time_add" yaml:"real_time_add" gorm:"column:real_time_add"`
}

// TableName TableName
func (Target) TableName() string {
	return "basic_targets"
}

// InsertTarget InsertTarget
func (c *DBAgent) InsertTarget(record *Target) error {
	err := c.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&record).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// InsertMultiTarget InsertMultiTarget
func (c *DBAgent) InsertMultiTarget(recordArr []*Target) error {
	err := c.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.CreateInBatches(&recordArr, multiInsertBatchSize).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteAllTarget DeleteAllTarget
func (c *DBAgent) DeleteAllTarget() error {
	err := c.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Not("id = 0").Unscoped().Delete(&Target{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteMultiTargetByDate DeleteMultiTargetByDate
func (c *DBAgent) DeleteMultiTargetByDate(date time.Time) error {
	err := c.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("trade_day = ?", date).Unscoped().Delete(&Target{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// GetTargetsByDate GetTargetsByDate
func (c *DBAgent) GetTargetsByDate(date time.Time) (targetArr []*Target, err error) {
	result := c.DB.Preload("Stock").Where("trade_day = ?", date).Find(&targetArr)
	return targetArr, result.Error
}

// CheckExistTargetsByDateStockID CheckExistTargetsByDateStockID
func (c *DBAgent) CheckExistTargetsByDateStockID(date time.Time, stockID int64) (bool, error) {
	var count int64
	err := c.DB.Model(&Target{}).Where("trade_day = ? and stock_id = ?", date, stockID).Count(&count).Error
	if count > 0 {
		return true, err
	}
	return false, err
}
