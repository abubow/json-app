package models

type Post struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Author    string `json:"author"`
	Published bool   `json:"published"`
}

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

// type Post struct {
// 	ID       uint      `json:"id" gorm:"primary_key"`
// 	UserID   uint      `json:"user_id" gorm:"not null"`
// 	Image    string    `json:"image"`
// 	Caption  string    `json:"caption"`
// 	Likes    int       `json:"likes"`
// 	Comments []Comment `json:"comments" gorm:"foreignkey:PostID"`
// }

// type Comment struct {
// 	ID      uint   `json:"id" gorm:"primary_key"`
// 	UserID  uint   `json:"user_id" gorm:"not null"`
// 	PostID  uint   `json:"post_id" gorm:"not null"`
// 	Content string `json:"content"`
// }
// type StaticImageData struct {
// 	// fields and types equivalent to StaticImageData in TypeScript

// 	URL string `json:"url"`
// 	Alt string `json:"alt"`
// }

// type ProductSummary struct {
// 	ID            string  `json:"id"`
// 	Title         string  `json:"title"`
// 	Price         float64 `json:"price"`
// 	PreviousPrice float64 `json:"previousPrice"`
// 	Reviews       struct {
// 		Count  int `json:"count"`
// 		Rating int `json:"rating"`
// 	} `json:"reviews"`
// 	Image StaticImageData `json:"image"`
// }

// type CollectionType struct {
// 	CollectionTitle string            `json:"collectionTitle"`
// 	Collection      []*ProductSummary `json:"collection"`
// }

// type ProductData struct {
// 	ProductSummary
// 	Instructions   string   `json:"instructions"`
// 	DeliveryDetail string   `json:"deliveryDetail"`
// 	ProductDetails string   `json:"productDetails"`
// 	Style          []string `json:"style"`
// 	Size           []string `json:"size"`
// 	ReviewPosts    []struct {
// 		Name   string `json:"name"`
// 		Rating int    `json:"rating"`
// 		Review string `json:"review"`
// 	} `json:"reviewPosts"`
// 	Images          []StaticImageData `json:"images"`
// 	Remaining       int               `json:"remaining"`
// 	CollectionTitle string            `json:"collectionTitle"`
// }
