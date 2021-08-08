package models

import (
	"github.com/Isaiah-peter/netfley-backend/pkg/config"
	"github.com/jinzhu/gorm"
)

var ()

type List struct {
	gorm.Model
	Title   string `gorm:"unique"`
	Type    string
	Genre   string
	Content []MyContent
}

type MyContent struct{}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&List{})
}
