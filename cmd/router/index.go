package router

import (
	"fmt"
	"gohub/cmd/controller"
	"net/http"
)

func SetRouter(){
  fmt.Printf("SetRouter\n")
	http.HandleFunc("/v1/api/users", controller.FetchAllUser)
	http.HandleFunc("/v1/api/user", controller.FetchUser)
}
