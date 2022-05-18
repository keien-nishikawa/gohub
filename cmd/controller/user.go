package controller

import (
	"fmt"
	db "gohub/cmd/db"
	"gohub/cmd/model"
	"log"
	"net/http"
)


func FetchAllUser(w http.ResponseWriter, r *http.Request){
	var username string  = model.Username()
	fmt.Fprintf(w, "fetchAll controller")
	log.Println("fetchAllUser:" + username)
	db.ConnectionDatabase()
}

func FetchUser(w http.ResponseWriter, r *http.Request){
	var username string  = model.Username()
	fmt.Fprintf(w, "fetchUser controller")
	log.Println("fetchUser" + username)
}
