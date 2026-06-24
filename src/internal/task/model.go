package task

import "time"

type Task struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UpdateTask struct {
	ID          uint
	Title       string
	Description string
	Completed   string
}
