package model

import "time"

type UserInput struct {
	ID        int 		`gorm:"primaryKey"`
	UserID    int
	Content   string
	Length    int
	Height    int
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User 		`gorm:"foreignKey:UserID"`
}