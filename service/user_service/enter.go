package userservice

import "goblog/models"

type UserService struct {
	userModel models.UserModel
}

func NewUserService(user models.UserModel) *UserService {
	return &UserService{
		userModel: user,
	}
}
