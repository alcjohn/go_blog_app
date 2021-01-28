package core

import (
	"fmt"

	"github.com/goava/di"
	"github.com/jinzhu/gorm"
)

type Controller struct {
	di.Inject
	DB *gorm.DB
}

func (c Controller) Render() {
	fmt.Println("coucou")
}
