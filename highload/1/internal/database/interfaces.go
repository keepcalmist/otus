package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/keepcalmist/otus/highload/1/internal/database/models"
	"github.com/keepcalmist/otus/highload/1/internal/database/postgres"
	_ "github.com/lib/pq"
)

type UsersRepo interface {
	Register(ctx context.Context, req *models.UserRegisterReq) (*models.UserRegisterResp, error)
	Get(ctx context.Context, req *models.GetUserFilter) ([]*models.User, error)
	InsertForCmd(ctx context.Context, tableName string, values map[string]interface{}) error
}

type Repo struct {
	Users UsersRepo
}

func NewRepo(ctx context.Context, database string) (*Repo, error) {
	db, err := sql.Open("postgres", database)
	if err != nil {
		return nil, fmt.Errorf("opennig db error:%w", err)
	}
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping error:%w", err)
	}

	usersRepo, err := postgres.NewUserRepo(ctx, db)
	if err != nil {
		return nil, fmt.Errorf("creating userRepo err:%w", err)
	}
	return &Repo{
		Users: usersRepo,
	}, nil
}
