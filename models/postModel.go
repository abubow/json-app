package models

// type User struct {
// 	ID           uint   `json:"id" gorm:"primary_key"`
// 	Username     string `json:"username" gorm:"unique;not null"`
// 	Email        string `json:"email" gorm:"unique;not null"`
// 	Password     string `json:"password" gorm:"not null"`
// 	ProfileImage string `json:"profile_image"`
// 	Followers    []User `json:"followers" gorm:"many2many:followers"`
// 	Followings   []User `json:"followings" gorm:"many2many:followings"`
// 	Posts        []Post `json:"posts" gorm:"foreignkey:UserID"`
// }

type Post struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Published bool   `json:"published"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Author    User   `json:"author" gorm:"foreignkey:AuthorID"`
	AuthorID  uint   `json:"author_id"`
}
