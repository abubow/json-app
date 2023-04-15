package models

import "time"

type User struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	Username     string    `json:"username" gorm:"unique;not null"`
	Email        string    `json:"email" gorm:"unique;not null"`
	Password     string    `json:"password" gorm:"not null"`
	ProfileImage string    `json:"profile_image"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Posts        []Post    `json:"posts" gorm:"foreignkey:UserID"`
	DummyFlag    bool      `json:"dummy_flag"`
}

type UserInfo struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	Username     string    `json:"username" gorm:"unique;not null"`
	Email        string    `json:"email" gorm:"unique;not null"`
	ProfileImage string    `json:"profile_image"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Posts        []Post    `json:"posts" gorm:"foreignkey:UserID"`
	DummyFlag    bool      `json:"dummy_flag"`
}
