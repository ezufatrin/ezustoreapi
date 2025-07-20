package infrastructure

import (
	"ezustore/src/domain"
	"ezustore/src/infrastructure/database"
)

type userRepo struct{}

func NewUserRepo(db interface{}) domain.UserRepository { // db param kept for symmetry; not used directly
	return &userRepo{}
}

func (r *userRepo) Register(user *domain.User) error {
	return database.DB.Create(user).Error
}

func (r *userRepo) Update(user *domain.User) error {
	return database.DB.Save(user).Error
}

func (r *userRepo) GetByID(id uint) (*domain.User, error) {
	var user domain.User
	err := database.DB.First(&user, id).Error
	return &user, err
}

func (r *userRepo) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
