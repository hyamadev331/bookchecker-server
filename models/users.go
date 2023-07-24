package models

import (
	"time"

	"gorm.io/gorm"
)

var database *gorm.DB

type User struct {
	id        int
	name      string
	age       int
	height    float32
	weight    float32
	createdAt time.Time
	updatedAt time.Time
}

// func init() {
// 	database = db.connectDB()
// }

// func (u *User) createUser() *User {
// 	database.Create(&u)
// }
