package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Isaiah-peter/netfley-backend/pkg/models"
	"github.com/Isaiah-peter/netfley-backend/pkg/utils"
	"github.com/gorilla/mux"
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

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	var movie = &models.Movies{}
	token := utils.UseToken(r)
	isAdmin := token["IsAdmin"]

	if isAdmin == true {
		vars := mux.Vars(r)
		moviesID := vars["id"]

		ID, err := strconv.ParseInt(moviesID, 0, 0)
		if err != nil {
			panic(err)
		}
		utils.ParseBody(r, movie)
		movieDetail, db := models.GetMovieByIdModel(ID)
		db.Save(&movieDetail)
	}

	fmt.Println("you are not a admin")
}

func GetMovieById(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(r)

	vars := mux.Vars(r)
	moviesID := vars["id"]

	ID, err := strconv.ParseInt(moviesID, 0, 0)

	if err != nil {
		panic(err)
	}

	movieDetail, _ := models.GetMovieByIdModel(ID)
	res, _ := json.Marshal(movieDetail)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	token := utils.UseToken(r)

	isAdmin := token["IsAdmin"]

	if isAdmin == true {
		vars := mux.Vars(r)
		moviesID := vars["id"]

		ID, err := strconv.ParseInt(moviesID, 0, 0)
		if err != nil {
			panic(err)
		}
		models.DeleteMovieModel(ID)
		res, _ := json.Marshal("the movie have been deleted")
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}

	if isAdmin == false {
		res, _ := json.Marshal("you are not a admin")
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func GetRandomMovie(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(r)
	Movie := models.GetMovieModel()
	res, _ := json.Marshal(Movie)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
