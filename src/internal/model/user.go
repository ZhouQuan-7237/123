package model

import "time"

type User struct {
	ID           int    	 `gorm:"primaryKey"`
	Name         string 	 `gorm:"type:varchar(100)"`
	Email        string 	 `gorm:"type:varchar(255)"`
	PasswordHash string 	 `gorm:"type:varchar(255)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Inputs       []UserInput `gorm:"foreignKey:UserID"`
}



