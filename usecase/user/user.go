package user

import (
	"todoListApp/entities"
	"todoListApp/repository/user"
)

type UserUseCase struct {
	UserRepository user.UserRepositoryInterface
}

func NewUserUseCase(userRepo user.UserRepositoryInterface) UserUseCaseInterface {
	return &UserUseCase{
		UserRepository: userRepo,
	}

}

func (uuc *UserUseCase) GetAll() ([]entities.User, error) {
	users, err := uuc.UserRepository.GetAll()
	return users, err
}

func (uuc *UserUseCase) GetUserById(id int) (entities.User, error) {
	user, err := uuc.UserRepository.GetUserById(id)
	return user, err
}

func (uuc *UserUseCase) CreateUser(user entities.User) error {
	err := uuc.UserRepository.CreateUser(user)
	return err
}

func (uuc *UserUseCase) DeleteUser(id int) error {
	err := uuc.UserRepository.DeleteUser(id)
	return err
}

func (uuc *UserUseCase) UpdateUser(id int, user entities.User) error {
	err := uuc.UserRepository.UpdateUser(id, user)
	return err
}
