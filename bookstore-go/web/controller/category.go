package controller

import (
	"bookstore/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	CategoryService *service.CategoryService
}

func NewCategoryController() *CategoryController {
	return &CategoryController{
		CategoryService: service.NewCategoryService(),
	}
}

func (c *CategoryController) GetCategoryList(ctx *gin.Context) {
	categories, err := c.CategoryService.GetCategoryList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":  -1,
			"msg":   "获取分类列表失败",
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "获取分类列表成功",
		"data": categories,
	})
}
