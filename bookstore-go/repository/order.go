package repository

import (
	"bookstore/global"
	"bookstore/model"

	"gorm.io/gorm"
)

type OrderDAO struct {
	db *gorm.DB
}

func NewOrderDAO() *OrderDAO {
	return &OrderDAO{
		db: global.GetDB(),
	}
}

func (o *OrderDAO) CreateOrderWithItems(order *model.Order, items []*model.OrderItem) error {
	return o.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			return err
		}
		for _, item := range items {
			item.OrderID = order.ID
			if err := tx.Create(item).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
