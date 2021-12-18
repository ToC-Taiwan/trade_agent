// Package dbagent package dbagent
package dbagent

import (
	"time"

	"gorm.io/gorm"
)

// OrderStatus OrderStatus
type OrderStatus struct {
	gorm.Model `json:"-" swaggerignore:"true"`

	Stock   *Stock `json:"stock,omitempty" yaml:"stock" gorm:"foreignKey:StockID"`
	StockID int64  `json:"stock_id,omitempty" yaml:"stock_id" gorm:"column:stock_id"`

	Action    int64     `json:"action,omitempty" yaml:"action" gorm:"column:action"`
	Price     float64   `json:"price,omitempty" yaml:"price" gorm:"column:price"`
	Quantity  int64     `json:"quantity,omitempty" yaml:"quantity" gorm:"column:quantity"`
	Status    int64     `json:"status,omitempty" yaml:"status" gorm:"column:status"`
	OrderID   string    `json:"order_id,omitempty" yaml:"order_id" gorm:"column:order_id"`
	OrderTime time.Time `json:"order_time,omitempty" yaml:"order_time" gorm:"column:order_time"`
}

// ActionListMap ActionListMap
var ActionListMap = map[string]int64{
	"Buy":  1,
	"Sell": 2,
}

// StatusListMap StatusListMap
var StatusListMap = map[string]int64{
	"PendingSubmit": 1, // 傳送中
	"PreSubmitted":  2, // 預約單
	"Submitted":     3, // 傳送成功
	"Failed":        4, // 失敗
	"Canceled":      5, // 已刪除
	"Filled":        6, // 完全成交
	"Filling":       7, // 部分成交
}

// InsertOrderStatus InsertOrderStatus
func (c *DBAgent) InsertOrderStatus(record *OrderStatus) error {
	err := c.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&record).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// InsertOrUpdateMultiOrderStatus InsertOrUpdateMultiOrderStatus
func (c *DBAgent) InsertOrUpdateMultiOrderStatus(recordArr []*OrderStatus) error {
	err := c.DB.Transaction(func(tx *gorm.DB) error {
		for _, v := range recordArr {
			var dbRecord OrderStatus
			result := tx.Model(&OrderStatus{}).Where("order_id = ?", v.OrderID).First(&dbRecord)
			if result.RowsAffected == 0 {
				err := c.InsertOrderStatus(v)
				if err != nil {
					return err
				}
			} else {
				tmp := v
				tmp.Model = dbRecord.Model
				if err := tx.Save(&tmp).Error; err != nil {
					return err
				}
			}
		}
		return nil
	})
	return err
}

// DeleteAllOrderStatus DeleteAllOrderStatus
func (c *DBAgent) DeleteAllOrderStatus() error {
	err := c.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Not("id = 0").Unscoped().Delete(&OrderStatus{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}