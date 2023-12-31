// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"Clean/Mongo-Crud/pkg/api"
	"Clean/Mongo-Crud/pkg/api/handler"
	"Clean/Mongo-Crud/pkg/db"
	"Clean/Mongo-Crud/pkg/repository"
	"Clean/Mongo-Crud/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI() (*http.ServerHTTP, error) {
	client, err := db.ConnectionDB()
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRepository(client)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)
	serverHTTP := http.NewServerHTTP(userHandler)
	return serverHTTP, nil
}
