package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password []byte `json:"."`
}

// SetPassword sets the password for the User.
//
// Parameters:
// - password: the password to be set.
//
// Returns: None.
func (user *User) SetPassword(password string) {
	hashed_password, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user.Password = hashed_password
}

func (user *User) PasswordCompare(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
