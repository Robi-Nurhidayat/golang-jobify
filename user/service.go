package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(input RegisterInput) (User, error)
	Login(input LoginInput) (User, error)
	GetUserById(id int) (User, error)
	GetAllUser() ([]User, error)
}

type userService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}

func (service *userService) Register(input RegisterInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.LastName = input.LastName
	user.Location = input.Location

	if len(input.Avatar) != 0 {
		user.Avatar = input.Avatar
	}
	user.Avatar = "avatar.jpg"
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		return user, err
	}
	user.Password = string(hashPassword)
	users, err := service.GetAllUser()

	if err != nil {
		return user, err
	}

	if len(users) > 0 {
		user.Role = "user"
	} else {
		user.Role = "admin"
	}

	userRegister, err := service.repository.Register(user)
	if err != nil {
		return userRegister, err
	}

	return userRegister, nil
}

func (service *userService) Login(input LoginInput) (User, error) {

	email := input.Email
	password := input.Password

	findEmail, err := service.repository.FindByEmail(email)
	if err != nil {
		return findEmail, err
	}

	if findEmail.Id == 0 {
		return findEmail, errors.New("no user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(findEmail.Password), []byte(password))

	if err != nil {
		return findEmail, err
	}

	return findEmail, nil
}

func (service *userService) GetUserById(id int) (User, error) {

	user, err := service.repository.FindById(id)

	if err != nil {
		return user, err
	}

	if user.Id == 0 {
		return user, errors.New("No user found on whith that ID")
	}

	return user, nil
}

func (service *userService) GetAllUser() ([]User, error) {

	users, err := service.repository.GetAllUser()

	if err != nil {
		return users, err
	}

	return users, nil
}
