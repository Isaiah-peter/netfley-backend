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
	Content []MyContent `gorm:"many2many:content;"`
}

type MyContent struct {
	ID string
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&List{})
}

func (m *List) CreateList() *List {
	db.NewRecord(m)
	db.Create(m)
	return m
}

func GetListModel() []List {
	var List []List
	db.Find(&List)
	return List
}

func GetListByIdModel(Id int64) (*List, *gorm.DB) {
	var getList List
	db.Where("ID=?", Id).Find(&getList)
	return &getList, db
}

func DeleteListModel(Id int64) List {
	var List List
	db.Where("ID=?", Id).Delete(List)
	return List
}
