package controllers

import (
	"encoding/json"
	"net/http"
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
		Name:    user.Username,
		Email:   user.Email,
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
