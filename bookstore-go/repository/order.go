package repository

import (
	"bookstore/global"
	"bookstore/model"

	"errors"

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

func (o *OrderDAO) GetOrderByID(orderID int) (*model.Order, error) {
	var order model.Order
	if err := o.db.Preload("OrderItems.Book").First(&order, orderID).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (o *OrderDAO) UpdateOrderStatus(order *model.Order) error {
	// 订单号，订单状态->1, 销量+1，库存-1
	return o.db.Debug().Transaction(func(tx *gorm.DB) error {
		for _, item := range order.OrderItems {
			var book model.Book
			if err := tx.First(&book, item.BookID).Error; err != nil {
				return errors.New("图书不存在")
			}
			if book.Stock < item.Quantity {
				return errors.New("库存不足")
			}
		}
		if tx.Model(&model.Order{}).Where("id = ?", order.ID).Update("status", 1).Updates(
			map[string]interface{}{
				"status":       1,
				"is_paid":      1,
				"payment_time": gorm.Expr("NOW()"),
			},
		).Error != nil {
			return errors.New("更新订单状态失败")
		}

		for _, item := range order.OrderItems {
			if err := tx.Model(&model.Book{}).Where("id = ?", item.BookID).Updates(
				map[string]interface{}{
					"stock": gorm.Expr("stock - ?", item.Quantity),
					"sale":  gorm.Expr("sale + ?", item.Quantity),
				},
			).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
