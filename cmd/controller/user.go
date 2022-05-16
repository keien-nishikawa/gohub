package controller

import (
	"fmt"
	db "gohub/cmd/db"
	"gohub/cmd/model"
	"log"
	"net/http"
)


func FetchAllUser(w http.ResponseWriter, r *http.Request){
	log.Print("fetchAll users")
  fmt.Fprintf(w, "fetchAll controller")
	model.Username()
	db.ConnectionDatabase()
}

func FetchUser(w http.ResponseWriter, r *http.Request){
	log.Print("fetch user")
  fmt.Fprintf(w, "fetchAll controller")
	model.Username()
}
