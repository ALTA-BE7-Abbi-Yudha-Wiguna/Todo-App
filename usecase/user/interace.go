package user

import "todoListApp/entities"

type UserUseCaseInterface interface {
	GetAll() ([]entities.User, error)
	GetUserById(id int) (entities.User, error)
	CreateUser(user entities.User) error
	DeleteUser(id int) error
	UpdateUser(id int, user entities.User) error
}
