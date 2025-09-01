package repository

import (
	"bookstore/global"
	"bookstore/model"

	"gorm.io/gorm"
)

type FavoriteDAO struct {
	db *gorm.DB
}

func NewFavoriteDAO() *FavoriteDAO {
	return &FavoriteDAO{
		db: global.GetDB(),
	}
}

func (f *FavoriteDAO) AddFavorite(userID, bookID int) error {
	return f.db.Debug().Create(&model.Favorite{
		UserID: userID,
		BookID: bookID,
	}).Error
}

func (f *FavoriteDAO) DeleteFavorite(userID, bookID int) error {
	return f.db.Debug().Where("user_id = ? AND book_id = ?", userID, bookID).Delete(&model.Favorite{}).Error
}

func (f *FavoriteDAO) GetUserFavorites(userID int, page, pageSize int, timeFilter string) ([]*model.Favorite, error) {
	var favs []*model.Favorite

	if err := f.db.Debug().Preload("Book").Where("user_id = ?", userID).Find(&favs).Error; err != nil {
		return nil, err
	}

	return favs, nil
}
