package models

import (
	"github.com/Isaiah-peter/netfley-backend/pkg/config"
	"github.com/Isaiah-peter/netfley-backend/pkg/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

type User struct {
	gorm.Model
	Username   string `gorm:"unique" binding:"required"`
	Email      string `gorm:"unique" form:"email" binding:"required"`
	Password   string `binding:"required,min=6"`
	ProfilePic *string
	IsAdmin    bool `gorm:"default:false"`
}

type Token struct {
	UserID  int
	IsAdmin bool
	jwt.StandardClaims
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) NewUser() *User {
	hashpassword, err := utils.HashPassword(u.Password)
	if err != nil {
		panic(err)
	}
	u.Password = hashpassword
	db.NewRecord(u)
	db.Create(u)
	return u
}

func GetUser() []User {
	var User []User
	db.Find(&User).Limit(10)
	return User
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("ID=?", Id).Find(&getUser)
	return &getUser, db
}

func DeleteUser(Id int64) User {
	var user User
	db.Where("ID=?", Id).Delete(user)
	return user
}
