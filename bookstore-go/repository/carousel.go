package repository

import (
	"bookstore/global"
	"bookstore/model"
	"gorm.io/gorm"
)

type CarouselDAO struct {
	db *gorm.DB
}

func NewCarouselDAO() *CarouselDAO {
	return &CarouselDAO{
		db: global.GetDB(),
	}
}

func (c *CarouselDAO) GetCarouselList() ([]model.Carousel, error) {
	var carousel []model.Carousel
	if err := c.db.Debug().Model(&model.Carousel{}).Find(&carousel).Error; err != nil {
		return nil, err
	}
	return carousel, nil
}
