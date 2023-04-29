package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  []byte `json:"-"`
}

/**
* Hash input string for user password
*/
func (user *User) SetPassword(password string) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.Password = hashed
}
/**
* Compare the users hashed password to input string
*/
func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
