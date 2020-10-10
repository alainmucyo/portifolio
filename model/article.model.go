package model

import (
	"errors"
	"github.com/alainmucyo/my_brand/config"
	"time"
)

type Article struct {
	ID           uint64    `gorm:"primary_key,auto_increment" json:"id"`
	Title        string    `json:"title" validate:"required,min=3"`
	Image        string    `json:"image"`
	Content      string    `json:"content" validate:"required,min=3"`
	Likes        int64     `json:"likes" gorm:"default:0"`
	CommentCount int32     `json:"comment_count" gorm:"default:0"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (Article) Get() ([]Article, error) {
	articles := make([]Article, 0)
	if config.Database.Find(&articles).Error != nil {
		return nil, errors.New("Error while selecting")
	}
	return articles, nil

}

func (article Article) Create() (Article, error) {
	ctx := config.Database.Begin()
	if ctx.Error != nil {
		return article, errors.New("Error start")
	}
	if config.Database.Save(&article).Error != nil {
		ctx.Rollback()
		return article, errors.New("Error save")
	}
	if ctx.Commit().Error != nil {
		ctx.Rollback()
		return article, errors.New("Error commit")
	}
	return article, nil
}

func (Article) FindById(articleId int) (Article, error) {
	var article Article
	if config.Database.Where("id = ?", articleId).Take(&article).Error != nil {
		return Article{}, errors.New("Article not found")
	}
	return article, nil
}

func (article Article) Update(articleId uint64) (Article, error) {
	var updatedArticle Article
	if config.Database.Model(&updatedArticle).Where("id = ?", articleId).Updates(article).Error != nil {
		return Article{}, errors.New("Article not found")
	}
	return updatedArticle, nil
}
func (Article) Delete(articleId uint64) error {
	if config.Database.Delete(&Article{}, articleId).Error != nil {
		return errors.New("Unable to delete!")
	}
	return nil
}
