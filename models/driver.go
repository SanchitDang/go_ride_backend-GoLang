// models/Driver.go
package models

type Driver struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func NewDriver(name string, email string, phone string) *Driver {
	return &Driver{name, email, phone}
}
