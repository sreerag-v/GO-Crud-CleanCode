
//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"Clean/Mongo-Crud/pkg/api"
	"Clean/Mongo-Crud/pkg/api/handler"
	"Clean/Mongo-Crud/pkg/db"
	"Clean/Mongo-Crud/pkg/usecase"
	"Clean/Mongo-Crud/pkg/repository"
)


func InitializeAPI() (*http.ServerHTTP, error) {
	wire.Build(
		db.ConnectionDB,

		 repository.NewUserRepository,

		 usecase.NewUserUseCase,

		 handler.NewUserHandler,

		 http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
