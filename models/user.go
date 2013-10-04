package models

import "code.google.com/p/go.crypto/bcrypt"

type User struct {
	Email string
	PasswordHash []byte
}

func (user *User) HasPassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password))
}

func NewUser(email string, password string) *User {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil
	}
	return &User{email, hashedPassword}
}
