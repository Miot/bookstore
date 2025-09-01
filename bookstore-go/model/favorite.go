package model

import "time"

type Favorite struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	UserID    int       `gorm:"not null" json:"user_id"`
	BookID    int       `gorm:"not null" json:"book_id"`
	CreatedAt time.Time `json:"created_at"`
	Book      *Book     `gorm:"foreignKey:BookID" json:"book,omitempty"`
}

func (f *Favorite) TableName() string {
	return "favorites"
}
