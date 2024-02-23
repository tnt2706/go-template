package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	Name        *string   `json:"name" bson:"usr_name,omitempty"`
	Email       *string   `json:"email" bson:"usr_email,omitempty"`
	Phone       *string   `json:"phone" bson:"usr_phone,omitempty"`
	Sex         *string   `json:"sex" bson:"usr_sex,omitempty"`
	Avatar      *string   `json:"avatar" bson:"usr_avatar omitempty"`
	DateOfBirth *string   `json:"date_of_birth" bson:"usr_date_of_birth,omitempty"`
	Roles       []*string `json:"roles" bson:"usr_roles,omitempty"`
	Status      *string   `json:"status" bson:"usr_status,omitempty"`
	Password    string    `json:"password" bson:"user_password,omitempty"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"updatedAt,omitempty"`
}

func (u *User) HashPassword(password string) error {
	bytePassword := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(passwordHash)
	return nil
}

func (u *User) ComparePassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func (u *User) Add() {
	u.UpdatedAt = time.Now()
	u.CreatedAt = time.Now()
}

func (u *User) Update() {
	u.UpdatedAt = time.Now()
}
