package controller

import (
	"bookstore/service"

	"github.com/gin-gonic/gin"
	"net/http"
)

type CarouselController struct {
	CarouselService *service.CarouselService
}

func NewCarouselController() *CarouselController {
	return &CarouselController{
		CarouselService: service.NewCarouselService(),
	}
}

func (c *CarouselController) GetCarouselList(ctx *gin.Context) {
	carousel, err := c.CarouselService.GetCarouselList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":  -1,
			"msg":   "获取轮播图列表失败",
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "获取轮播图列表成功",
		"data": carousel,
	})
}
