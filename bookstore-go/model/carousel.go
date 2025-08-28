package model

import "time"

// 轮播图
type Carousel struct {
	ID          int       `json:"id" gorm:"primarykey"`
	Title       string    `json:"title" gorm:"not null;comment:轮播图标题"`
	Description string    `json:"description" gorm:"type:text;comment:轮播图描述"`
	ImageURL    string    `json:"image_url" gorm:"not null;comment:轮播图图片URL"`
	LinkURL     string    `json:"link_url" gorm:"comment:轮播图链接URL"`
	SortOrder   int       `json:"sort_order" gorm:"default:0;comment:排序"`
	IsActive    bool      `json:"is_active" gorm:"default:true;comment:是否激活"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (c *Carousel) TableName() string {
	return "carousels"
}
