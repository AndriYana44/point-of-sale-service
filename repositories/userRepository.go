package repositories

import (
	"gin-learn/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repository *UserRepository) FindByEmail(
	email string,
) (*models.User, error) {
	var user models.User

	err := repository.db.
		Where("email = ?", email).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *UserRepository) FindByID(
	id uint,
) (*models.User, error) {
	var user models.User

	err := repository.db.
		First(&user, id).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
