package service

import (
	"bookstore/model"
	"bookstore/repository"
)

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

func (f *FavoriteService) DeleteFavorite(userID int, bookID int) error {
	return f.FavoriteDB.DeleteFavorite(userID, bookID)
}

func (f *FavoriteService) GetFavoriteList(userID int, page, pageSize int, timeFilter string) ([]*model.Favorite, int64, error) {
	fav, err := f.FavoriteDB.GetUserFavorites(userID, page, pageSize, timeFilter)
	if err != nil {
		return nil, 0, err
	}
	total := len(fav)
	start := (page - 1) * pageSize
	end := start + pageSize
	if start >= total {
		return []*model.Favorite{}, 0, nil
	}
	if end > total {
		end = total
	}

	return fav[start:end], int64(total), nil
}

func (f *FavoriteService) CheckFavorite(userID int, bookID int) (bool, error) {
	return f.FavoriteDB.CheckFavorite(userID, bookID)
}

func (f *FavoriteService) GetFavoriteCount(userID int) (int64, error) {
	return f.FavoriteDB.GetFavoriteCount(userID)
}
