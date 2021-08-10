package repository

import (
	"go-ddd/domain/entity"
	"gorm.io/gorm"
)

type IUser interface {
	Create(db *gorm.DB, user *entity.User) (uint, error)
	GetByEmail(db *gorm.DB, email string) (*entity.User, error)
	GetByRecoveryToken(db *gorm.DB, recoveryToken string) (*entity.User, error)
	Update(db *gorm.DB, user *entity.User) error

	EmailExists(db *gorm.DB, email string) (bool, error)
}
