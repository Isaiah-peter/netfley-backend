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

//CREATE

func (m *Movies) CreateMovie() *Movies {
	db.NewRecord(m)
	db.Create(m)
	return m
}

func GetMovieModel() []Movies {
	var movie []Movies
	db.Find(&movie)
	return movie
}

func GetMovieByIdModel(Id int64) (*Movies, *gorm.DB) {
	var getMovie Movies
	db.Where("ID=?", Id).Find(&getMovie)
	return &getMovie, db
}

func DeleteMovieModel(Id int64) Movies {
	var movie Movies
	db.Where("ID=?", Id).Delete(movie)
	return movie
}

func GetMovieWhereTypeIsIsSeries(series bool) interface{} {
	var movies []Movies
	tvSeries := db.Where("is_series = ?", series).Find(&movies).Value
	return tvSeries
}
