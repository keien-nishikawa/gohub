package router

import (
	"fmt"
	"gohub/cmd/controller"
	"net/http"
)

func SetRouter(){
	http.HandleFunc("/v1/api/users", controller.FetchAllUser)
	http.HandleFunc("/v1/api/user", controller.FetchUser)
	printRouterInfo() // consoleにローティング情報を表示
}

func printRouterInfo(){
	fmt.Printf("Routeing info\n")
	fmt.Printf("http://localhost:8080\n")
	fmt.Printf("http://localhost:8080/v1/api/users\n")
	fmt.Printf("http://localhost:8080/v1/api/user\n")
}
