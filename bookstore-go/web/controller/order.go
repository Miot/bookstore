package controller

import (
	"bookstore/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	OrderService *service.OrderService
}

func NewOrderController() *OrderController {
	return &OrderController{
		OrderService: service.NewOrderService(),
	}
}

func (o *OrderController) CreateOrder(c *gin.Context) {
	var req service.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  -1,
			"msg":   "请求参数错误",
			"error": err.Error(),
		})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": -1,
			"msg":  "用户未登录",
		})
		return
	}

	req.UserID = userID.(int)
	order, orderErr := o.OrderService.CreateOrder(&req)
	if orderErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  -1,
			"msg":   "创建订单失败",
			"error": orderErr.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "创建订单成功",
		"data": order,
	})
}
