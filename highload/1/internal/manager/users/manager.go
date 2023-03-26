package users

import (
	"context"
	"crypto/sha1"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/keepcalmist/otus/highload/1/internal/database"
	"github.com/keepcalmist/otus/highload/1/internal/database/models"
	"github.com/keepcalmist/otus/highload/1/internal/manager"
	"time"
)

type man struct {
	repo *database.Repo

	secretKey      []byte
	expireDuration time.Duration
}

func New(repo *database.Repo) UserService {
	return &man{repo: repo}
}

func (m *man) Register(ctx context.Context, req *UserRegisterReq) (*UserRegisterResp, error) {
	req.Password = m.hashPass(req.Password)
	return m.saveUserToDB(ctx, req)
}

func (m *man) saveUserToDB(ctx context.Context, req *UserRegisterReq) (*UserRegisterResp, error) {
	resp, err := m.repo.Users.Register(ctx, &models.UserRegisterReq{
		FirstName:  req.FirstName,
		SecondName: req.SecondName,
		Age:        req.Age,
		Biography:  req.Biography,
		City:       req.City,
		Password:   req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &UserRegisterResp{UserID: resp.UserID}, nil
}

func (m *man) GetByID(ctx context.Context, id int64) (*User, error) {
	users, err := m.repo.Users.Get(ctx, &models.GetUserFilter{UserIDs: []int64{id}})
	if err != nil {
		return nil, err
	}

	if len(users) < 1 {
		return nil, manager.NotFound
	}

	return castUsers(users)[0], nil
}

func (m *man) GetByFilter(ctx context.Context, filter *SearchFilter) ([]*User, error) {
	if filter == nil {
		return nil, manager.BadRequest
	}

	users, err := m.repo.Users.Get(ctx, &models.GetUserFilter{FirstName: filter.FirstName, SecondName: filter.SecondName})
	if err != nil {
		return nil, err
	}

	if len(users) < 1 {
		return nil, manager.NotFound
	}

	return castUsers(users), nil
}

func (m *man) Login(ctx context.Context, req *Login) (string, error) {
	users, err := m.repo.Users.Get(ctx, &models.GetUserFilter{UserIDs: []int64{req.ID}})
	if err != nil {
		return "", err
	}

	if len(users) < 1 {
		return "", manager.NotFound
	}

	if string(m.hashPass(req.Password)) != string(users[0].Password) {
		return "", manager.IncorrectPassOrLogin
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(m.expireDuration)),
		},
		id: req.ID,
	})

	return token.SignedString(m.secretKey)
}

func (m *man) hashPass(pass []byte) []byte {
	hash := sha1.New()
	hash.Write(pass)
	hash.Write(m.secretKey)
	return hash.Sum(nil)
}
