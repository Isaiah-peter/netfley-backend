package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Isaiah-peter/netfley-backend/pkg/models"
	"github.com/Isaiah-peter/netfley-backend/pkg/utils"
)

var (
	Movie models.Movies
)

//create
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	movie := &models.Movies{}
	token := utils.UseToken(r)

	IsAdmin := token["IsAdmin"]

	if IsAdmin == true {
		utils.ParseBody(r, movie)
		u := movie.CreateMovie()
		res, _ := json.Marshal(u)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}

	fmt.Println("you are not a admin")

}
