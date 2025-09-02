package repository

import (
	"bookstore/global"
	"bookstore/model"

	"gorm.io/gorm"
)

type CategoryDAO struct {
	db *gorm.DB
}

func NewCategoryDAO() *CategoryDAO {
	return &CategoryDAO{
		db: global.GetDB(),
	}
}

func (c *CategoryDAO) GetCategoryList() ([]*model.Category, error) {
	var categories []*model.Category
	if err := c.db.Debug().Model(&model.Category{}).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
