package service

import (
	"bookstore/model"
	"bookstore/repository"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type OrderService struct {
	OrderDB *repository.OrderDAO
	BookDB  *repository.BookDAO
}
type CreateOrderRequest struct {
	UserID int          `json:"user_id"`
	Items  []OrderItems `json:"items"`
}
type OrderItems struct {
	BookID   int `json:"book_id"`
	Quantity int `json:"quantity"`
	Price    int `json:"price"`
}

func NewOrderService() *OrderService {
	return &OrderService{
		OrderDB: repository.NewOrderDAO(),
		BookDB:  repository.NewBookDAO(),
	}
}

func (o *OrderService) CreateOrder(req *CreateOrderRequest) (*model.Order, error) {
	if len(req.Items) == 0 {
		return nil, errors.New("订单项不能为空")
	}

	// 库存校验
	err := o.CheckStockAvailability(req.Items)
	if err != nil {
		return nil, err
	}
	// 生成订单号
	orderNo := fmt.Sprintf("ORD%d%06d", time.Now().UnixNano(), rand.Intn(1000000))
	var totalAmount int
	var orderItems []*model.OrderItem
	for _, item := range req.Items {
		subtotal := item.Quantity * item.Price
		orderItems = append(orderItems, &model.OrderItem{
			BookID:   item.BookID,
			Quantity: item.Quantity,
			Price:    item.Price,
			Subtotal: subtotal,
		})
		totalAmount += subtotal
	}
	// 支付
	order := &model.Order{
		OrderNo:     orderNo,
		UserID:      req.UserID,
		TotalAmount: totalAmount,
		Status:      0,
		IsPaid:      false,
	}
	createErr := o.OrderDB.CreateOrderWithItems(order, orderItems)
	if createErr != nil {
		return nil, createErr
	}
	return order, nil
}

func (o *OrderService) CheckStockAvailability(items []OrderItems) error {
	for _, item := range items {
		book, err := o.BookDB.GetBookByID(item.BookID)
		if err != nil {
			return errors.New("图书不存在")
		}
		if book.Status != 1 {
			return errors.New("图书已下架")
		}
		if book.Stock < item.Quantity {
			return errors.New("库存不足")
		}
	}
	return nil
}

func (o *OrderService) GetOrderList(userID int, page, pageSize int) ([]*model.Order, int64, error) {
	return o.OrderDB.GetOrderList(userID, page, pageSize)
}

func (o *OrderService) PayOrder(orderID int) error {
	// 检查订单是否存在
	order, err := o.OrderDB.GetOrderByID(orderID)
	if err != nil {
		return err
	}
	if order.Status != 0 {
		return errors.New("订单状态异常")
	}
	// 更新订单状态
	return o.OrderDB.UpdateOrderStatus(order)
}
