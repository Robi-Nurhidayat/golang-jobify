package user

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
	panic("asda")
}

func (service *userService) Login(input LoginInput) (User, error) {
	//TODO implement me
	panic("implement me")
}
