package user

import "golang.org/x/crypto/bcrypt"

type UserService interface {
	Register(input RegisterInput) (User, error)
	Login(input LoginInput) (User, error)
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
	user.Role = "user"
	if len(input.Avatar) != 0 {
		user.Avatar = input.Avatar
	}

	user.Avatar = "avatar.jpg"

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}

	user.Password = string(hashPassword)

	userRegister, err := service.repository.Register(user)
	if err != nil {
		return userRegister, err
	}

	return userRegister, nil
}

func (service *userService) Login(input LoginInput) (User, error) {
	//TODO implement me
	panic("implement me")
}
