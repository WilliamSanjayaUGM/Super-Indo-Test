package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"size:64,index"`
	Password string `gorm:"size:255"`
}

func CheckUserAvailability(email string) bool {
	var user User
	DB.Where("email = ?", email).First(&user)
	return user.ID == 0 // true if the email is available
}

func UserCreate(email string, password string) *User {
	// Password Hashing
	// hashedPassword, err := bcry.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// if err != nil {
	// 	return nil
	// }
	entry := User{Email: email, Password: password}
	DB.Create(&entry)
	return &entry
}

func UserMatchPassword(email string, password string) *User {
	var user User
	DB.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		return &user
	}

	// err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	// if err != nil {
	// 	return &User{}
	// } else {
	// 	return &user
	// }

	if password == user.Password {
		return &user
	} else {
		return &User{}
	}
}

func UserFromId(id uint) *User {
	var user User
	DB.Where("id = ?", id).First(&user)
	return &user
}
