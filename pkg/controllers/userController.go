package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Isaiah-peter/netfley-backend/pkg/config"
	"github.com/Isaiah-peter/netfley-backend/pkg/models"
	"github.com/Isaiah-peter/netfley-backend/pkg/utils"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var (
	db   = config.GetDB()
	User models.User
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	newuser := models.GetUser()
	res, _ := json.Marshal(newuser)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	newUser := &models.User{}
	utils.ParseBody(r, newUser)
	u := newUser.NewUser()
	res, _ := json.Marshal(u)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Login(w http.ResponseWriter, r *http.Request) {
	newUser := &models.User{}
	err := json.NewDecoder(r.Body).Decode(newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	u := FindOne(newUser.Email, newUser.Password)
	json.NewEncoder(w).Encode(u)

}

func FindOne(email string, password string) map[string]interface{} {
	user := &models.User{}

	if err := db.Where("Email = ?", email).First(user).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp

	}

	expireAt := time.Now().Add(time.Minute * 100000).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		return resp
	}

	tk := &models.Token{
		UserID:  int(user.ID),
		IsAdmin: user.IsAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)

	tokenString, err := token.SignedString([]byte("my_secret_key"))
	if err != nil {
		panic(err)
	}

	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = tokenString
	resp["user"] = user
	return resp
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user = &models.User{}
	utils.ParseBody(r, user)
	token, err := utils.VerifyToken(r)
	if err != nil {
		panic(err)
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		panic(ok)
	}

	ID, err := strconv.ParseInt(fmt.Sprintf("%.f", claim["UserID"]), 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	fmt.Println(ID)
	userDetail, rd := models.GetUserById(ID)
	if user.Username != "" {
		userDetail.Username = user.Username
		fmt.Println("userDetailfmt:", userDetail.Username)
	}
	if user.Email != "" {
		userDetail.Email = user.Email
	}
	if user.Password != "" {
		hashpassword, err := utils.HashPassword(user.Password)
		if err != nil {
			panic(err)
		}
		userDetail.Password = hashpassword
	}
	if user.ProfilePic != nil {
		userDetail.ProfilePic = user.ProfilePic
	}
	rd.Save(&userDetail)
	fmt.Println("userDetailfmt:", userDetail.Username)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	token, err := utils.VerifyToken(r)
	if err != nil {
		fmt.Println("error in verifying token or expire token")
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("token expired")
	}

	ID, err := strconv.ParseInt(fmt.Sprintf("%.f", claim["UserID"]), 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
		log.Fatalln(err)
	}
	userDetail, _ := models.GetUserById(ID)
	res, _ := json.Marshal(userDetail)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	token, err := utils.VerifyToken(r)
	if err != nil {
		fmt.Println("error in verifying token or expire token")
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("token expired")
	}

	ID, err := strconv.ParseInt(fmt.Sprintf("%.f", claim["UserID"]), 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
		log.Fatalln(err)
	}
	user := models.DeleteUser(ID)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
