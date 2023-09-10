package handler

import (
	handlerInterface "Clean/Mongo-Crud/pkg/api/handler/interface"
	"Clean/Mongo-Crud/pkg/domain"
	services "Clean/Mongo-Crud/pkg/usecase/interface"
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type UserHandler struct {
	UserUseCase services.UserUseCase
}

func NewUserHandler(usecase services.UserUseCase) handlerInterface.UserHandler {
	return &UserHandler{
		UserUseCase: usecase,
	}
}

func (usr UserHandler) CreateUser(rs http.ResponseWriter, re *http.Request, _ httprouter.Params) {
	user := domain.Users{}

	err := json.NewDecoder(re.Body).Decode(&user)

	if err != nil {
		rs.WriteHeader(http.StatusBadRequest)
	}
	err = usr.UserUseCase.CreateUser(re.Context(), user)

	if err != nil {
		rs.WriteHeader(http.StatusInternalServerError)

	}
	userJson, err := json.Marshal(user)
	if err != nil {
		rs.WriteHeader(http.StatusInternalServerError)
		return
	}

	rs.Header().Set("Content-Type", "application/json")
	rs.WriteHeader(http.StatusOK) // Use 200 for success
	rs.Write(userJson)
}

func (usr *UserHandler) GetUserByid(rs http.ResponseWriter, re *http.Request, p httprouter.Params) {
	Uid := p.ByName("id")

	User, err := usr.UserUseCase.GetUserByid(re.Context(), Uid)
	if err != nil {
		rs.WriteHeader(http.StatusInternalServerError)
	}

	Ujson, err := json.Marshal(User)

	if err != nil {
		rs.WriteHeader(http.StatusInternalServerError) // Handle JSON marshaling error
		return
	}

	rs.Header().Set("Content-Type", "application/json")
	rs.WriteHeader(http.StatusOK) // Use 200 for success
	rs.Write(Ujson)
}

func (usr *UserHandler) UpdateUserById(rs http.ResponseWriter, re *http.Request, p httprouter.Params) {
	Uid := p.ByName("id")

	User := domain.Users{}
	err := json.NewDecoder(re.Body).Decode(&User)

	if err != nil {
		rs.WriteHeader(http.StatusBadRequest)
	}

	err = usr.UserUseCase.UpdateUserById(re.Context(), Uid, User)
	if err != nil {
		rs.WriteHeader(http.StatusInternalServerError)
	}

	rs.Header().Set("Content-Type", "application/json")
	rs.WriteHeader(http.StatusOK) // Use 200 for success

	response := map[string]string{"status": "User Successfully Updated"}
	JsonResponse, _ := json.Marshal(response)
	rs.Write(JsonResponse)
}

func (usr *UserHandler) DeleteUserById(rs http.ResponseWriter, re *http.Request, p httprouter.Params) {
	Uid := p.ByName("id")

	err := usr.UserUseCase.DeleteUserById(re.Context(), Uid)
	if err != nil {
		rs.WriteHeader(http.StatusInternalServerError)
	}

	rs.Header().Set("Content-Type", "application/json")
	rs.WriteHeader(http.StatusOK) // Use 200 for success

	response := map[string]string{"status": "User Successfully Deleted"}
	JsonResponse, _ := json.Marshal(response)
	rs.Write(JsonResponse)
}
