package domain

import "context"

type User struct {
	Name     string `json:"Name"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type UserUsecase interface {
	Signup(c context.Context, user User) error
	Delete(c context.Context, user User) error
	Login(c context.Context, user User) error
}

type UserRepository interface {
	CreateUser(c context.Context, user User) error
	ReadUser(c context.Context, user User) error
	UpdateUser(c context.Context, user User) error
	DeleteUser(c context.Context, user User) error
}
