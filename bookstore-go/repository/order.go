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

func (o *OrderDAO) GetOrderList(userID int, page, pageSize int) ([]*model.Order, int64, error) {
	var orders []*model.Order
	var total int64
	if err := o.db.Debug().Model(&model.Order{}).Where("user_id = ?", userID).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	offset := (page - 1) * pageSize
	if err := o.db.Preload("OrderItems.Book").Where("user_id = ?", userID).Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&orders).Error; err != nil {
		return nil, 0, err
	}
	return orders, total, nil
}
