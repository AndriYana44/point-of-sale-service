package repositories

import (
	"gin-learn/models"

	"gorm.io/gorm"
)

type RefreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{
		db: db,
	}
}

func (repository *RefreshTokenRepository) Create(
	refreshToken *models.RefreshToken,
) error {
	return repository.db.Create(refreshToken).Error
}

func (repository *RefreshTokenRepository) FindByTokenHash(
	tokenHash string,
) (*models.RefreshToken, error) {
	var refreshToken models.RefreshToken

	err := repository.db.
		Where("token_hash = ?", tokenHash).
		First(&refreshToken).
		Error

	if err != nil {
		return nil, err
	}

	return &refreshToken, nil
}

func (repository *RefreshTokenRepository) Revoke(
	id uint,
) error {
	return repository.db.
		Model(&models.RefreshToken{}).
		Where("id = ?", id).
		Update("revoked_at", gorm.Expr("NOW()")).
		Error
}
