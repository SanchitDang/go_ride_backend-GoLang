// models/User.go
package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func NewUser(name string, email string, phone string) *User {
	return &User{}
}
