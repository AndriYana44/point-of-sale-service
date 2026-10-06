package models

import "time"

type RefreshToken struct {
	ID        uint       `gorm:"column:id;primaryKey"`
	UserID    uint       `gorm:"column:user_id"`
	TokenHash string     `gorm:"column:token_hash"`
	ExpiresAt time.Time  `gorm:"column:expires_at"`
	RevokedAt *time.Time `gorm:"column:revoked_at"`
	CreatedAt time.Time  `gorm:"column:created_at"`
}

func (RefreshToken) TableName() string {
	return "refresh_tokens"
}
