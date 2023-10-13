package user

import "gorm.io/gorm"
type UserRepository interface {
	Register(user User) (User, error)
}

type userRepositoryImpl struct {
	db *gorm.DB
}


func (repository *userRepositoryImpl) Register(user User) (User,error) {
	err := repository.db.Create(&user).Error

	if err != nil {
		return user,nil
	}

	return user,nil

}
