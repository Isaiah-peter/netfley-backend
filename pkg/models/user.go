package models

import (
	"github.com/Isaiah-peter/netfley-backend/pkg/config"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

type User struct {
	gorm.Model
	Username   string `gorm:"unique"`
	Email      string `gorm:"unique"`
	Password   string
	ProfilePic *string
	IsAdmin    bool `gorm:"default:false"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}
