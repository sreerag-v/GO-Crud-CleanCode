package interfaces

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserHandler interface{
	CreateUser(rs http.ResponseWriter, re *http.Request, _ httprouter.Params)
	GetUserByid(rs http.ResponseWriter, re *http.Request, p httprouter.Params) 
	UpdateUserById(rs http.ResponseWriter, re *http.Request, p httprouter.Params) 
	DeleteUserById(rs http.ResponseWriter, re *http.Request, p httprouter.Params)
}