package models

import (
	"github.com/Isaiah-peter/netfley-backend/pkg/config"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

type Movies struct {
	gorm.Model
	Title    string `gorm:"unique"`
	Disc     string
	Img      string
	ImgTitle string
	ImgSm    string
	trailer  string
	Video    string
	Year     string
	Limit    int
	genre    string
	IsSeries bool `gorm:"default:false"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Movies{})
}
