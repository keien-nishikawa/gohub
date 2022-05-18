package main

import (
	router "gohub/cmd/config"
	"gohub/cmd/util"
	"log"
	"net/http"
)

func main() {
	router.SetRouter()
 	err := http.ListenAndServe(":8080",nil)
	util.LoggingSettings("develop.log")
	log.Println("test start.")
	if err != nil {
		log.Fatal(err)
	}
}

