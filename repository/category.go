package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"
	"errors"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error)
	StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error)
	StoreManyCategory(ctx context.Context, categories []entity.Category) error
	GetCategoryByID(ctx context.Context, id int) (entity.Category, error)
	UpdateCategory(ctx context.Context, category *entity.Category) error
	DeleteCategory(ctx context.Context, id int) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error) {
	getId := []entity.Category{}
	err := r.db.WithContext(ctx).Model(&entity.Category{}).Where("user_id = ?", id).Find(&getId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []entity.Category{}, nil
	}
	if err != nil {
		return nil, err
	}
	return getId, nil
}

func (r *categoryRepository) StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error) {
	store := r.db.WithContext(ctx).Model(&entity.Category{}).Create(category)
	err = store.Error
	if err != nil {
		return 0, err
	}
	categoryId = category.ID
	return categoryId, nil
}

func (r *categoryRepository) StoreManyCategory(ctx context.Context, categories []entity.Category) error {
	return r.db.WithContext(ctx).Model(&entity.Category{}).CreateInBatches(&categories, len(categories)).Error
}

func (r *categoryRepository) GetCategoryByID(ctx context.Context, id int) (entity.Category, error) {
	var getCatId entity.Category
	err := r.db.WithContext(ctx).Model(&entity.Category{}).Where("ID = ?", id).Find(&getCatId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.Category{}, nil
	}
	if err != nil {
		return entity.Category{}, err
	}
	return getCatId, nil
}

func (r *categoryRepository) UpdateCategory(ctx context.Context, category *entity.Category) error {
	return r.db.WithContext(ctx).Model(&entity.Category{}).Where("ID = ?", category.ID).Updates(category).Error
}

func (r *categoryRepository) DeleteCategory(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Model(&entity.Category{}).Where("ID = ?", id).Delete(&entity.Category{}).Error
}
