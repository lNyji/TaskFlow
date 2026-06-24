package task

import (
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func (r *Repository) Order(tasks []Task) any {
	panic("unimplemented")
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(task Task) error {
	return r.db.Create(&task).Error
}

func (r *Repository) List() ([]Task, error) {
	var tasks []Task

	return tasks, r.db.Order("id asc").Find(&tasks).Error
}

func (r *Repository) Update(input UpdateTask) error {
	var task Task

	if err := r.db.First(&task, input.ID).Error; err != nil {
		return err
	}

	if input.Title != "" {
		task.Title = input.Title
	}

	if input.Description != "" {
		task.Description = input.Description
	}

	switch input.Completed {
	case "s":
		task.Completed = true
	case "n":
		task.Completed = false
	}

	task.UpdatedAt = time.Now()

	return r.db.Save(&task).Error
}

func (r *Repository) Delete(id uint) error {
	var task Task

	return r.db.First(&task, id).Delete(&task, id).Error
}
