package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  []byte `json:"-"`
}

func (user *User) SetPassword(password string) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.Password = hashed
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
