package model

import (
	"encoding/json"

	"golang.org/x/crypto/bcrypt"
)

type UserId int64

type User struct {
	ID       UserId `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Users []*User

func (u *User) JsonMarshal() ([]byte, error) {
	resp, err := json.MarshalIndent(u, "", "\t\t")
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *User) CryptPassword() error {
	encodedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return err
	}
	u.Password = string(encodedPassword)
	return nil
}

func (us *Users) JsonMarshal() ([]byte, error) {
	resp, err := json.MarshalIndent(us, "", "\t\t")
	if err != nil {
		return nil, err
	}
	return resp, nil
}
