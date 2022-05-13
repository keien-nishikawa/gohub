package main

import (
	"fmt"
	"gohub/router"
	"log"
	"net/http"
)

func main() {
  fmt.Printf("Hello world\n")
	router.SetRouter()
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Fatal(err)
	}
}
