package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"
	"errors"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTasks(ctx context.Context, id int) ([]entity.Task, error)
	StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error)
	GetTaskByID(ctx context.Context, id int) (entity.Task, error)
	GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error)
	UpdateTask(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, id int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) GetTasks(ctx context.Context, id int) ([]entity.Task, error) {
	var getTasks []entity.Task
	err := r.db.WithContext(ctx).Model(&entity.Task{}).Where("user_id = ?", id).Find(&getTasks).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []entity.Task{}, nil
	}
	if err != nil {
		return nil, err
	}
	return getTasks, nil
}

func (r *taskRepository) StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error) {
	store := r.db.WithContext(ctx).Model(&entity.Task{}).Create(task)
	err = store.Error
	if err != nil {
		return 0, err
	}
	taskId = task.ID
	return taskId, nil
}

func (r *taskRepository) GetTaskByID(ctx context.Context, id int) (entity.Task, error) {
	var task entity.Task
	err := r.db.WithContext(ctx).Model(&entity.Task{}).Where("ID = ?", id).First(&task).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.Task{}, nil
	}
	if err != nil {
		return entity.Task{}, err
	}
	return task, nil
}

func (r *taskRepository) GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error) {
	var getCategory []entity.Task
	err := r.db.WithContext(ctx).Model(&entity.Task{}).Where("category_id = ?", catId).Find(&getCategory).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []entity.Task{}, nil
	}
	if err != nil {
		return nil, err
	}
	return getCategory, nil
}

func (r *taskRepository) UpdateTask(ctx context.Context, task *entity.Task) error {
	return r.db.WithContext(ctx).Model(&entity.Task{}).Where("ID = ?", task.ID).Updates(task).Error
}

func (r *taskRepository) DeleteTask(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Model(&entity.Task{}).Where("ID = ?", id).Delete(&entity.Task{}).Error
}
