package router

import (
	"gohub/controller"
	"net/http"
)

func UserRouter(){
	http.HandleFunc("v1/api/users", controller.FetchAllUser)
	http.HandleFunc("v1/api/user", controller.FetchUser)
}
