package main

import (
	router "gohub/cmd/config"
	"log"
	"net/http"
)

func main() {
	router.SetRouter()
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Fatal(err)
	}
	router.ConnectionDatabase()
}
