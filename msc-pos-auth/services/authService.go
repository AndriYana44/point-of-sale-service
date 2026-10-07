package services

import (
	"errors"
	"time"

	"gin-learn/helpers"
	"gin-learn/models"
	"gin-learn/repositories"

	"gorm.io/gorm"
)

type AuthService struct {
	userRepository         *repositories.UserRepository
	refreshTokenRepository *repositories.RefreshTokenRepository
}

func NewAuthService(
	userRepository *repositories.UserRepository,
	refreshTokenRepository *repositories.RefreshTokenRepository,
) *AuthService {
	return &AuthService{
		userRepository:         userRepository,
		refreshTokenRepository: refreshTokenRepository,
	}
}

func (service *AuthService) Login(
	email string,
	password string,
) (string, string, error) {

	user, err := service.userRepository.FindByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", "", errors.New("invalid email or password")
		}

		return "", "", err
	}

	if !helpers.CheckPassword(password, user.Password) {
		return "", "", errors.New("invalid email or password")
	}

	accessToken, err := helpers.GenerateAccessToken(
		user.ID,
		user.Email,
	)

	if err != nil {
		return "", "", err
	}

	refreshToken, err := helpers.GenerateRefreshToken()
	if err != nil {
		return "", "", err
	}

	refreshTokenModel := &models.RefreshToken{
		UserID:    user.ID,
		TokenHash: helpers.HashRefreshToken(refreshToken),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	if err := service.refreshTokenRepository.Create(
		refreshTokenModel,
	); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (service *AuthService) GetUserByID(
	id uint,
) (*models.User, error) {
	return service.userRepository.FindByID(id)
}

func (service *AuthService) RefreshToken(
	refreshToken string,
) (string, error) {

	tokenHash := helpers.HashRefreshToken(refreshToken)

	storedToken, err := service.refreshTokenRepository.FindByTokenHash(
		tokenHash,
	)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("invalid refresh token")
		}

		return "", err
	}

	if storedToken.RevokedAt != nil {
		return "", errors.New("refresh token has been revoked")
	}

	if time.Now().After(storedToken.ExpiresAt) {
		return "", errors.New("refresh token has expired")
	}

	user, err := service.userRepository.FindByID(
		storedToken.UserID,
	)

	if err != nil {
		return "", err
	}

	accessToken, err := helpers.GenerateAccessToken(
		user.ID,
		user.Email,
	)

	if err != nil {
		return "", err
	}

	return accessToken, nil
}
