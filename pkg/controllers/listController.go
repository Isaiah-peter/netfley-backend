package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Isaiah-peter/netfley-backend/pkg/models"
	"github.com/Isaiah-peter/netfley-backend/pkg/utils"
)

var (
	List models.List
)

func CreateList(w http.ResponseWriter, r *http.Request) {
	list := &models.List{}
	token := utils.UseToken(r)

	IsAdmin := token["IsAdmin"]

	log.Fatalln(IsAdmin)

	if IsAdmin == true {
		utils.ParseBody(r, list)
		u := list.CreateList()
		res, _ := json.Marshal(u)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(res)
	}

	fmt.Println("you are not a admin")

}
