package models

import "time"

type Post struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	Published bool       `json:"published"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Author    User       `json:"author" gorm:"foreignkey:AuthorID"`
	AuthorID  uint       `json:"author_id"`
	Likes     []*User    `json:"likes" gorm:"many2many:likes"`
	Comments  []*Comment `json:"comments" gorm:"foreignkey:PostID"`
	UserID    uint       `json:"-"`
}

type Comment struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Author    User      `json:"author" gorm:"foreignkey:AuthorID"`
	AuthorID  uint      `json:"author_id"`
	Post      *Post     `json:"post" gorm:"foreignkey:PostID"`
	PostID    uint      `json:"post_id"`
	Likes     []*User   `json:"likes" gorm:"many2many:likes"`
	// may have other comment as children or may not
	Comments []*Comment `json:"comments" gorm:"foreignkey:ParentID"`
	UserID   uint       `json:"-"`
	ParentID uint       `json:"-"`
}
