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

func (b *BookDAO) GetBooksByPage(page, pageSize int) ([]model.Book, int64, error) {
	var books []model.Book
	var total int64
	if err := b.db.Model(&model.Book{}).Debug().Where("status = ?", 1).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := b.db.Debug().Where("status = ?", 1).Offset(offset).Limit(pageSize).Find(&books).Error; err != nil {
		return nil, 0, err
	}

	return books, total, nil
}

func (b *BookDAO) SearchBooksWithPage(keyword string, page, pageSize int) ([]model.Book, int64, error) {
	var books []model.Book
	var total int64
	searchCondition := b.db.Debug().Where("status = ? AND (title LIKE ? OR author LIKE ? OR description LIKE ?)", 1, "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")

	err := searchCondition.Model(&model.Book{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := searchCondition.Offset(offset).Limit(pageSize).Find(&books).Error; err != nil {
		return nil, 0, err
	}

	return books, total, nil
}

func (b *BookDAO) GetBookDetail(id int) (model.Book, error) {
	var book model.Book
	if err := b.db.Debug().Where("status = ?", 1).First(&book, id).Error; err != nil {
		return book, err
	}
	return book, nil
}

func (b *BookDAO) GetBookByID(id int) (*model.Book, error) {
	var book model.Book
	if err := b.db.Debug().Where("status = ?", 1).First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}
