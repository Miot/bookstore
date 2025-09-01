package service

import "bookstore/repository"

type FavoriteService struct {
	FavoriteDB *repository.FavoriteDAO
}

func NewFavoriteService() *FavoriteService {
	return &FavoriteService{
		FavoriteDB: repository.NewFavoriteDAO(),
	}
}

func (f *FavoriteService) AddFavorite(userID int, bookID int) error {
	return f.FavoriteDB.AddFavorite(userID, bookID)
}
