package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title       string
	Description string
	Content     string
	UserID      int
	User        User
	PublishedAt time.Time
	IsPublished bool
}
