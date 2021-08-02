package models

import (
	"github.com/Isaiah-peter/netfley-backend/pkg/config"
	"github.com/jinzhu/gorm"
)

var ()

type Movies struct {
	gorm.Model
	Title    string `gorm:"unique"`
	Disc     string
	Img      string
	ImgTitle string
	ImgSm    string
	Trailer  string
	Video    string
	Year     string
	Limit    int
	Genre    string
	IsSeries bool `gorm:"default:false"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Movies{})
}
