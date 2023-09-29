package models

import (
	"time"
)

// User is a struct that represents user data in an application.
type User struct {
	Id        string    `gorm:"primaryKey;not null" json:"id"`
	Username  string    `gorm:"type:varchar(100);not null" json:"username"`
	Email     string    `gorm:"type:varchar(100);not null;unique;" json:"email"`
	Password  string    `gorm:"type:varchar(100);not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Photo     []Photo   `gorm:"foreignKey:userid;references:Id; constraint:On;Update:CASCADE;OnDelete:SET NULL;"`
}
