package main

import (
	router "gohub/cmd/config"
	"gohub/cmd/util"
	"log"
	"net/http"
)

func main() {
	// config logging
	util.LoggingSettings("./log/develop.log")
	// config routeing
	router.SetRouter()
	// error handling
 	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Fatal(err)
	}
}

