package repository

import (
	"bookstore/global"
	"bookstore/model"

	"gorm.io/gorm"
)

type BookDAO struct {
	db *gorm.DB
}

func NewBookDAO() *BookDAO {
	return &BookDAO{
		db: global.GetDB(),
	}
}

func (b *BookDAO) GetHotBooks(limit int) ([]model.Book, error) {
	var books []model.Book
	if err := b.db.Debug().Where("status = ?", 1).Order("sale DESC").Limit(limit).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (b *BookDAO) GetNewBooks(limit int) ([]model.Book, error) {
	var books []model.Book
	if err := b.db.Debug().Where("status = ?", 1).Order("created_at DESC").Limit(limit).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
