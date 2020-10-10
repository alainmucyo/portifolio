package model

import (
	"errors"
	"github.com/alainmucyo/my_brand/config"
	"html"
	"strings"
	"time"
)

type Query struct {
	ID        uint64    `gorm:"primary_key,auto_increment" json:"id"`
	Names     string    `json:"names" validate:"required,min=3"`
	Email     string    `json:"email" validate:"required,email"`
	Content   string    `json:"content" validate:"required,min=3"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (q *Query) Prepare() {
	q.ID = 0
	q.Names = html.EscapeString(strings.TrimSpace(q.Names))
	q.Email = html.EscapeString(strings.TrimSpace(q.Email))
	q.Content = html.EscapeString(strings.TrimSpace(q.Content))
	q.CreatedAt = time.Now()
	q.UpdatedAt = time.Now()
}

func (q Query) FindAll() ([]Query, error) {

	qus := make([]Query, 0)
	if config.Database.Find(&qus).Error != nil {
		return nil, errors.New("Error while getting queries")
	}
	return qus, nil

}

func (qury Query) Save() (Query, error) {
	ctx := config.Database.Begin()

	if ctx.Error != nil {
		return qury, errors.New("Error")
	}
	if config.Database.Save(&qury).Error != nil {
		ctx.Rollback()
		return qury, errors.New("Error")
	}
	if ctx.Commit().Error != nil {
		ctx.Rollback()
		return qury, errors.New("Error")
	}
	return qury, nil
}
