package user

type UserService interface {
	Register(user User) (User, error)
}

type userService struct {
	repository UserRepository
}

func (service *userService) Register()
