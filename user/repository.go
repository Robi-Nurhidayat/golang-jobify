package user

import "gorm.io/gorm"

type UserRepository interface {
	Register(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindById(id int) (User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (repository *userRepository) Register(user User) (User, error) {
	err := repository.db.Create(&user).Error

	if err != nil {
		return user, nil
	}

	return user, nil

}

func (repository *userRepository) FindByEmail(email string) (User, error) {
	var user User

	err := repository.db.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository *userRepository) FindById(id int) (User, error) {

	var user User

	err := repository.db.Where("id = ?", id).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
