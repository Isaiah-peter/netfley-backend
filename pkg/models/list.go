package models

import (
	"github.com/Isaiah-peter/netfley-backend/pkg/config"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

type List struct {
	gorm.Model
	Title   string `gorm:"unique"`
	Type    string
	genre   string
	Content []string
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&List{})
}
