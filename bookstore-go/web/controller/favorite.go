package controller

import (
	"bookstore/service"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

type FavoriteController struct {
	FavoriteService *service.FavoriteService
}

func NewFavoriteController() *FavoriteController {
	return &FavoriteController{
		FavoriteService: service.NewFavoriteService(),
	}
}

func getUserID(c *gin.Context) int {
	userID, exists := c.Get("userID")
	if !exists {
		return 0
	}
	return userID.(int)
}

func (f *FavoriteController) AddFavorite(c *gin.Context) {
	userID := getUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": -1,
			"msg":  "用户未登录",
		})
		return
	}

	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "书本ID格式错误",
		})
		return
	}

	err = f.FavoriteService.AddFavorite(userID, bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  "添加收藏失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "添加收藏成功",
	})
}

func (f *FavoriteController) DeleteFavorite(c *gin.Context) {
	userID := getUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": -1,
			"msg":  "用户未登录",
		})
		return
	}

	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "书本ID格式错误",
		})
		return
	}

	err = f.FavoriteService.DeleteFavorite(userID, bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  "删除收藏失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "删除收藏成功",
	})
}

func (f *FavoriteController) GetFavoriteList(c *gin.Context) {
	userID := getUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": -1,
			"msg":  "用户未登录",
		})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "12"))
	timeFilter := c.DefaultQuery("time_filter", "all")

	favs, total, err := f.FavoriteService.GetFavoriteList(userID, page, pageSize, timeFilter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  -1,
			"msg":   "获取收藏列表失败",
			"error": err.Error(),
		})
		return
	}

	totalPages := (total + int64(pageSize) - 1) / int64(pageSize)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "获取收藏列表成功",
		"data": gin.H{
			"favorites":    favs,
			"total":        total,
			"current_page": page,
			"total_pages":  totalPages,
		},
	})

}

func (f *FavoriteController) CheckFavorite(c *gin.Context) {
	userID := getUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": -1,
			"msg":  "用户未登录",
		})
		return
	}

	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "书本ID格式错误",
		})
		return
	}

	isFavorited, err := f.FavoriteService.CheckFavorite(userID, bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  "检查收藏状态失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "检查收藏状态成功",
		"data": gin.H{
			"is_favorited": isFavorited,
		},
	})
}

func (f *FavoriteController) GetFavoriteCount(c *gin.Context) {
	userID := getUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": -1,
			"msg":  "用户未登录",
		})
		return
	}

	count, err := f.FavoriteService.GetFavoriteCount(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  "获取收藏数量失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "获取收藏数量成功",
		"data": gin.H{
			"count": count,
		},
	})
}
