package model

import (
	"errors"
	"github.com/alainmucyo/my_brand/config"
	"time"
)

type Comment struct {
	ID        uint64 `gorm:"primary_key,auto_increment" json:"id"`
	Names     string `json:"names" validate:"required,min=3"`
	Content   string `json:"content" validate:"required,min=3"`
	ArticleID uint64 `json:"article_id" gorm:"column:article_id"`
	//Article   Article   `json:"article"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (comment Comment) AddComment() (Comment, error) {
	if config.Database.Save(&comment).Error != nil {
		return Comment{}, errors.New("Can't comment")
	}
	return comment, nil
}
