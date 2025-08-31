package controller

import (
	"bookstore/service"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	BookService *service.BookService
}

func NewBookController() *BookController {
	return &BookController{
		BookService: service.NewBookService(),
	}
}

func (b *BookController) GetHotBooks(c *gin.Context) {
	// 根据销量降序排序
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	books, err := b.BookService.GetHotBooks(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  -1,
			"msg":   "获取热销书籍失败",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "获取热销书籍成功",
		"data": books,
	})

}

func (b *BookController) GetNewBooks(c *gin.Context) {
	// 根据上架时间降序排序
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	books, err := b.BookService.GetNewBooks(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  -1,
			"msg":   "获取新品书籍失败",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "获取新品书籍成功",
		"data": books,
	})
}

func (b *BookController) GetBookList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "12"))

	books, total, err := b.BookService.GetBooksByPage(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  -1,
			"msg":   "获取书籍列表失败",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "获取书籍列表成功",
		"data": gin.H{
			"books":      books,
			"total":      total,
			"page":       page,
			"page_size":  pageSize,
			"total_size": (total + int64(pageSize) - 1) / int64(pageSize),
		},
	})
}

func (b *BookController) SearchBooks(c *gin.Context) {
	keyword := c.Query("q")
	if keyword == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "搜索关键词不能为空",
		})
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "12"))

	books, total, err := b.BookService.SearchBooksWithPage(keyword, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  -1,
			"msg":   "搜索书籍失败",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "搜索书籍成功",
		"data": gin.H{
			"books":      books,
			"total":      total,
			"page":       page,
			"page_size":  pageSize,
			"total_size": (total + int64(pageSize) - 1) / int64(pageSize),
		},
	})
}
