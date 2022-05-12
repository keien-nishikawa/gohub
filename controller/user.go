package controller

import (
	"fmt"
	"gohub/model"
	"log"
	"net/http"
)


func FetchAllUser(w http.ResponseWriter, r *http.Request){
	log.Print("fetchAll users")
  fmt.Fprintf(w, "fetchAll controller")
	model.Username()
}

func FetchUser(w http.ResponseWriter, r *http.Request){
	log.Print("fetchAll users")
  fmt.Fprintf(w, "fetchAll controller")
	model.Username()
}
