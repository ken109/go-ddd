package persistence

import (
	"github.com/pkg/errors"

	"go-gin-ddd/domain/entity"
	"go-gin-ddd/domain/repository"
	"go-gin-ddd/pkg/context"
)

type user struct{}

func NewUser() repository.IUser {
	return &user{}
}

func (u user) Create(ctx context.Context, user *entity.User) (uint, error) {
	db := ctx.DB()

	if err := db.Create(user).Error; err != nil {
		return 0, errors.WithStack(err)
	}
	return user.ID, nil
}

func (u user) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	db := ctx.DB()

	var dest entity.User
	err := db.Where(&entity.User{Email: email}).First(&dest).Error
	if err != nil {
		return nil, dbError(err)
	}
	return &dest, nil
}

func (u user) GetByRecoveryToken(ctx context.Context, recoveryToken string) (*entity.User, error) {
	db := ctx.DB()

	var dest entity.User
	err := db.Where(&entity.User{RecoveryToken: &recoveryToken}).First(&dest).Error
	if err != nil {
		return nil, dbError(err)
	}
	return &dest, nil
}

func (u user) Update(ctx context.Context, user *entity.User) error {
	db := ctx.DB()

	return db.Model(user).Updates(user).Error
}

func (u user) EmailExists(ctx context.Context, email string) (bool, error) {
	db := ctx.DB()

	var count int64
	if err := db.Model(&entity.User{}).Where(&entity.User{Email: email}).Count(&count).Error; err != nil {
		return false, errors.WithStack(err)
	}
	return count > 0, nil
}
