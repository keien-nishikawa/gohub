package main

import (
	router "gohub/cmd/config"
	"gohub/cmd/util"
	"log"
	"net/http"
	"os"
)

func main() {
	// config logging
	env_type := os.Getenv("ENV_TYPE")
	if env_type == "production" {
		util.LoggingSettings("./log/production.log")
	} else {
		util.LoggingSettings("./log/develop.log")
	}
	// config routeing
	router.SetRouter()
	// error handling
 	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Fatal(err)
	}
}

