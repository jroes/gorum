package models

import (
	"code.google.com/p/go.crypto/bcrypt"
	"bytes"
	"io/ioutil"
	"encoding/gob"
	"encoding/base64"
	"crypto/sha1"
	"os"
)

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

type UserStore interface {
	FindUser(string) (*User, error)
	SaveUser(User) error
}

type UserGobStore struct {
	Path string
}

func NewUserGobStore(path string) *UserGobStore {
	store := UserGobStore{path}
	err := os.MkdirAll(path, 0700)
	if err != nil {
		panic(err)
	}
	return &store
}

func (store UserGobStore) FindUser(email string) (*User, error) {
	emailSha := generateHash(email)
	userGob, err := ioutil.ReadFile(store.Path + emailSha + ".gob")
	if err != nil {
		return nil, err
	}

	userGobBuf := bytes.NewBuffer(userGob)
	decoder := gob.NewDecoder(userGobBuf)
	user := User{}
	err = decoder.Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (store UserGobStore) SaveUser(user User) error {
	emailSha := generateHash(user.Email)
	userGobBuf := new(bytes.Buffer)
	encoder := gob.NewEncoder(userGobBuf)
	encoder.Encode(user)
	return ioutil.WriteFile(store.Path + emailSha + ".gob", userGobBuf.Bytes(), 0600)
}

func generateHash(str string) string {
	hasher := sha1.New()
	hasher.Write([]byte(str))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
