package controller

import (
	"bookstore/service"

	"net/http"

	"strconv"

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

func (o *OrderController) GetOrderList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "12"))

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": -1,
			"msg":  "用户未登录",
		})
		return
	}

	orders, total, err := o.OrderService.GetOrderList(userID.(int), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  -1,
			"msg":   "获取订单列表失败",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "获取订单列表成功",
		"data": gin.H{
			"orders":     orders,
			"total":      total,
			"page":       page,
			"page_size":  pageSize,
			"total_page": (total + int64(pageSize) - 1) / int64(pageSize),
		},
	})
}

func (o *OrderController) PayOrder(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  "无效订单ID",
		})
		return
	}

	if err = o.OrderService.PayOrder(orderID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  -1,
			"msg":   "支付订单失败",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "支付订单成功",
	})
}
