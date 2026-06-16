package storage

import (
	"TaskFlow/internal/task"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("taskflow.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&task.Task{}); err != nil {
		return nil, err
	}

	return db, nil
}
