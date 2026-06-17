package note

import "time"

type Note struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	Title       string
	Description string
	Content     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
