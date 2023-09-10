package http

import (
	handlerInterface "Clean/Mongo-Crud/pkg/api/handler/interface"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type ServerHTTP struct {
	router *httprouter.Router
}

func NewServerHTTP(userHandler handlerInterface.UserHandler) *ServerHTTP {
	router := httprouter.New()

	//......Basic Crud only......//
	router.POST("/add-user", userHandler.CreateUser)
	router.GET("/get-user/:id", userHandler.GetUserByid)
	router.PUT("/update-user/:id", userHandler.UpdateUserById)
	router.DELETE("/delete-user/:id", userHandler.DeleteUserById)

	return &ServerHTTP{router: router}
}

func (sh *ServerHTTP) Start() {
	http.ListenAndServe(":8080", sh.router)
}
