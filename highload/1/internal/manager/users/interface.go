package users

import (
	"context"
	"github.com/dgrijalva/jwt-go/v4"
)

var _ UserService = &man{}

type UserService interface {
	Register(ctx context.Context, req *UserRegisterReq) (*UserRegisterResp, error)
	GetByID(ctx context.Context, id int64) (*User, error)
	Login(ctx context.Context, req *Login) (string, error)
}

type User struct {
	ID         int64
	FirstName  string
	SecondName string
	Age        int32
	Biography  string
	City       string
}

type UserRegisterReq struct {
	FirstName  string
	SecondName string
	Age        int32
	Biography  string
	City       string
	Password   []byte
}

type UserRegisterResp struct {
	UserID int64
}

type Claims struct {
	jwt.StandardClaims
	id int64
}

type Login struct {
	ID       int64
	Password []byte
}
