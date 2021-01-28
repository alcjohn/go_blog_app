package core

import (
	"github.com/goava/di"
	"github.com/jinzhu/gorm"
)

type IService interface {
}

type Service struct {
	di.Inject
	DB *gorm.DB
}
