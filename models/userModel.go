package models

type User struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	Username     string `json:"username" gorm:"unique;not null"`
	Email        string `json:"email" gorm:"unique;not null"`
	Password     string `json:"password" gorm:"not null"`
	ProfileImage string `json:"profile_image"`
	Followers    []User `json:"followers" gorm:"many2many:followers"`
	Followings   []User `json:"followings" gorm:"many2many:followings"`
}
