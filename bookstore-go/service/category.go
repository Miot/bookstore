package service

import (
	"bookstore/model"
	"bookstore/repository"
)

type CategoryService struct {
	CategoryDB *repository.CategoryDAO
}

func NewCategoryService() *CategoryService {
	return &CategoryService{
		CategoryDB: repository.NewCategoryDAO(),
	}
}

func (s *CategoryService) GetCategoryList() ([]*model.Category, error) {
	return s.CategoryDB.GetCategoryList()
}
