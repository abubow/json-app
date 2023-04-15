package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/a/json-app/initial"
	"github.com/a/json-app/models"
	"github.com/bxcodec/faker/v3"
)

// type User struct {
// 	ID           uint      `json:"id" gorm:"primary_key"`
// 	Username     string    `json:"username" gorm:"unique;not null"`
// 	Email        string    `json:"email" gorm:"unique;not null"`
// 	Password     string    `json:"password" gorm:"not null"`
// 	ProfileImage string    `json:"profile_image"`
// 	CreatedAt    time.Time `json:"created_at"`
// 	UpdatedAt    time.Time `json:"updated_at"`
// 	Posts        []Post    `json:"posts" gorm:"foreignkey:UserID"`
// }

// type Post struct {
// 	ID        uint       `json:"id" gorm:"primary_key"`
// 	Title     string     `json:"title"`
// 	Body      string     `json:"body"`
// 	Published bool       `json:"published"`
// 	CreatedAt time.Time  `json:"created_at"`
// 	UpdatedAt time.Time  `json:"updated_at"`
// 	Author    User       `json:"author" gorm:"foreignkey:AuthorID"`
// 	AuthorID  uint       `json:"author_id"`
// 	Likes     []*User    `json:"likes" gorm:"many2many:likes"`
// 	Comments  []*Comment `json:"comments" gorm:"foreignkey:PostID"`
// 	UserID    uint       `json:"-"`
// }

// type Comment struct {
// 	ID        uint       `json:"id" gorm:"primary_key"`
// 	Body      string     `json:"body"`
// 	CreatedAt time.Time  `json:"created_at"`
// 	UpdatedAt time.Time  `json:"updated_at"`
// 	Author    User       `json:"author" gorm:"foreignkey:AuthorID"`
// 	AuthorID  uint       `json:"author_id"`
// 	Post      *Post      `json:"post" gorm:"foreignkey:PostID"`
// 	PostID    uint       `json:"post_id"`
// 	Likes     []*User    `json:"likes" gorm:"many2many:likes"`
// 	Comments  []*Comment `json:"comments" gorm:"foreignkey:ParentID"`
// 	UserID    uint       `json:"-"`
// 	ParentID  uint       `json:"-"`
// }
// Create numComments comments for the above post

func createDummyData(numUsers, numPosts, numComments int) {
	// make text file to store user details or open it if it already exists
	f, err := os.OpenFile("dummyData.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	// make header for the text file if it is empty
	if fi, err := f.Stat(); err == nil && fi.Size() == 0 {
		l, err := f.WriteString("ID, Username, Email, Password, ProfileImage, CreatedAt, UpdatedAt" + "\n")
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}

		fmt.Println(l, "bytes written successfully")
	}
	// Create numUsers users
	for i := 0; i < numUsers; i++ {
		// make user details and store it in a file
		u := models.User{
			Username:     faker.TitleMale() + faker.Username(),
			Email:        faker.Email(),
			Password:     faker.Password(),
			ProfileImage: faker.URL(),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			DummyFlag:    true,
		}
		// store user details in a file
		l, err := f.WriteString(fmt.Sprintf("%v, %v, %v, %v, %v, %v, %v\n", u.ID, u.Username, u.Email, u.Password, u.ProfileImage, u.CreatedAt, u.UpdatedAt))
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
		fmt.Println(l, "bytes written successfully")
		// insert user into database
		err = initial.DB.Create(&u).Error
		if err != nil {
			log.Fatalf("Error creating user: %v", err)
		}

		// get the user from the database to get the ID
		val := initial.DB.Where("email = ?", u.Email).First(&u)
		if val.Error != nil {
			fmt.Println("Error getting user from database")
		}
		// print user details one by one in order id, username, email, password, profileimage, createdat, updatedat
		fmt.Println(u.ID, u.Username, u.Email, u.Password, u.ProfileImage, u.CreatedAt, u.UpdatedAt)

		// Create numPosts posts for each user
		for j := 0; j < numPosts; j++ {
			p := models.Post{
				Title:     faker.Sentence(),
				Body:      faker.Paragraph(),
				Published: true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				AuthorID:  u.ID,
				UserID:    0,
			}
			// insert post into user table as well as post table
			err := initial.DB.Model(&u).Association("Posts").Append(&p)
			if err != nil {
				log.Fatalf("Error creating post: %v", err)
			}

			// get the post from the database to get the ID
			val := initial.DB.Where("title = ?", p.Title).First(&p)
			if val.Error != nil {
				fmt.Println("Error getting post from database")
			}
			// for k := 0; k < numComments; k++ {
			// 	c := models.Comment{
			// 		Body:      faker.Paragraph(),
			// 		CreatedAt: time.Now(),
			// 		UpdatedAt: time.Now(),
			// 		AuthorID:  u.ID,
			// 		PostID:    p.ID,
			// 		UserID:    u.ID,
			// 		ParentID:  p.ID,
			// 		Comments:  nil,
			// 	}
			// 	// insert comment into comment table
			// 	err := initial.DB.Create(&c).Error
			// 	if err != nil {
			// 		log.Fatalf("Error creating comment: %v", err)
			// 	}

			// 	// insert comment into post table
			// 	err = initial.DB.Model(&p).Association("Comments").Append(&c)
			// 	if err != nil {
			// 		log.Fatalf("Error inserting comment: %v", err)
			// 	}

			// 	// get the comment from the database to get the ID
			// 	val := initial.DB.Where("body = ?", c.Body).First(&c)
			// 	if val.Error != nil {
			// 		fmt.Println("Error getting comment from database")
			// 	}
			// 	// print comment details one by one in order id, body, createdat, updatedat, authorid, postid, userid, parentid
			// 	fmt.Println(c.ID, c.Body, c.CreatedAt, c.UpdatedAt, c.AuthorID, c.PostID, c.UserID, c.ParentID)
			// }
		}
	}
	// close the file
	err = f.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	initial.LoadEnv()
	initial.ConnectToDB()
	initial.DB.AutoMigrate(&models.User{})
	initial.DB.AutoMigrate(&models.Post{})
	initial.DB.AutoMigrate(&models.Comment{})
	// make 20 posts, 5 comments for each post and
	// also create 3 users
	createDummyData(20, 5, 3)
	fmt.Println("Dummy data created")
}
