package model

import "time"

type Book struct {
	ID          int       `gorm:"primarykey" json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Price       int       `json:"price"`
	Discount    int       `json:"discount"`
	Type        string    `json:"type"`
	Stock       int       `json:"stock"`
	Status      int       `json:"status"` // 上架1，下架0
	Description string    `json:"description"`
	CoverUrl    string    `json:"cover_url"`
	ISBN        string    `json:"isbn"`
	Publisher   string    `json:"publisher"`
	Pages       int       `json:"pages"`
	Language    string    `json:"language"`
	Format      string    `json:"format"` // 装帧
	CategoryID  int       `json:"category_id"`
	Sale        int       `json:"sale"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (b *Book) TableName() string {
	return "books"
}
