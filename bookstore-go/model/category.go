package model

import "time"

type Category struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"not null;unique" json:"name"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
	Color       string    `json:"color"`
	Gradient    string    `json:"gradient"` // 渐变颜色
	Sort        int       `gorm:"default:0" json:"sort"`
	IsActive    bool      `gorm:"default:true" json:"is_active"`
	BookCount   int       `gorm:"default:0" json:"book_count"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (c *Category) TableName() string {
	return "categories"
}
