package interfaces

import (
	"Clean/Mongo-Crud/pkg/domain"
	"context"
)


type UserRepository interface{
	CreateUser(ctx context.Context,user domain.Users)error
	GetUserByid(ctx context.Context,Uid string)(domain.UsersResponse,error)
	UpdateUserById(ctx context.Context,Uid string,User domain.Users)error
	DeleteUserById(ctx context.Context,Uid string)error
}