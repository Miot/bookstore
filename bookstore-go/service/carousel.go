package service

import (
	"bookstore/model"
	"bookstore/repository"
)

type CarouselService struct {
	CarouselDB *repository.CarouselDAO
}

func NewCarouselService() *CarouselService {
	return &CarouselService{
		CarouselDB: repository.NewCarouselDAO(),
	}
}

func (c *CarouselService) GetCarouselList() ([]model.Carousel, error) {
	return c.CarouselDB.GetCarouselList()
}
