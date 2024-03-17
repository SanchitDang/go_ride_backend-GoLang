// models/User.go
package models

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func NewUser(name string, email string, phone string) *User {
	return &User{name, email, phone}
}
