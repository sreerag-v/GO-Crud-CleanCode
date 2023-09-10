package usecase

import (
	"Clean/Mongo-Crud/pkg/domain"
	interfaces "Clean/Mongo-Crud/pkg/repository/interface"
	usecaseInterface "Clean/Mongo-Crud/pkg/usecase/interface"
	"context"
)

type UserUseCase struct{
	userRepo interfaces.UserRepository
}

func NewUserUseCase(userRepo interfaces.UserRepository)usecaseInterface.UserUseCase{
	return &UserUseCase{
		userRepo:userRepo,
	}
}

func (usr *UserUseCase)	CreateUser(ctx context.Context,user domain.Users)error{
	return usr.userRepo.CreateUser(ctx,user)
}

func (usr *UserUseCase)	GetUserByid(ctx context.Context,Uid string)(domain.UsersResponse,error){
	return usr.userRepo.GetUserByid(ctx,Uid)
}

func (usr *UserUseCase)	UpdateUserById(ctx context.Context,Uid string,User domain.Users)error{
	return usr.userRepo.UpdateUserById(ctx,Uid,User)
}

func (usr *UserUseCase)	DeleteUserById(ctx context.Context,Uid string)error{
	return usr.userRepo.DeleteUserById(ctx,Uid)
}