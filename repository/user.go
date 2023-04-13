package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	getUser := entity.User{}
	err := r.db.WithContext(ctx).Model(&entity.User{}).Where("ID = ?", id).Find(&getUser).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.User{}, nil
	}
	if err != nil {
		return entity.User{}, err
	}
	return getUser, nil // TODO: replace this
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	getEmail := entity.User{}
	err := r.db.WithContext(ctx).Model(&entity.User{}).Where("email = ?", email).Find(&getEmail).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.User{}, nil
	}
	if err != nil {
		return entity.User{}, err
	}
	return getEmail, nil // TODO: replace this
}

func (r *userRepository) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	err := r.db.WithContext(ctx).Model(&entity.User{}).Create(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil // TODO: replace this
}

func (r *userRepository) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	err := r.db.WithContext(ctx).Model(&entity.User{}).Updates(user).Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", id).Delete(&entity.User{}).Error
}
