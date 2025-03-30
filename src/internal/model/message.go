package model

import "time"

type Message struct {
	ID        int 		`gorm:"primaryKey"`
	UserID    int
	Role      string 	`gorm:"type:varchar(20)"`
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User 		`gorm:"foreignKey:UserID"`
}