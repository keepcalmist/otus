package postgres

import (
	"context"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/keepcalmist/otus/highload/1/internal/database/models"
)

type userRepo struct {
	conn    *sql.DB
	builder *sq.StatementBuilderType
}

func NewUserRepo(ctx context.Context, db *sql.DB) (*userRepo, error) {
	b := sq.StatementBuilder.RunWith(db).PlaceholderFormat(sq.Dollar)
	repo := &userRepo{
		conn:    db,
		builder: &b,
	}

	return repo, nil
}

func (r *userRepo) Register(ctx context.Context, req *models.UserRegisterReq) (*models.UserRegisterResp, error) {
	qb := r.builder.Insert("users").Columns(
		"first_name",
		"second_name",
		"age",
		"biography",
		"city",
		"password",
	).Values(
		req.FirstName,
		req.SecondName,
		req.Age,
		req.Biography,
		req.City,
		req.Password,
	).Suffix(
		"RETURNING id")

	rows, err := qb.Query()
	if err != nil {
		return nil, fmt.Errorf("query err:%w", err)
	}

	var id int64
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return nil, fmt.Errorf("scan err:%w", err)
		}
	}

	return &models.UserRegisterResp{UserID: id}, err

}

func (r *userRepo) Get(ctx context.Context, req *models.GetUserFilter) ([]*models.User, error) {
	qb := r.builder.Select(
		"id",
		"first_name",
		"second_name",
		"age",
		"biography",
		"city",
		"password",
	).From("users").Where(
		sq.Eq{
			"id": req.UserIDs,
		},
	)
	fmt.Println(qb.ToSql())
	users := make([]*models.User, 0)
	rows, err := qb.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		user := new(models.User)
		err = rows.Scan(&user.ID, &user.FirstName, &user.SecondName, &user.Age, &user.Biography, &user.City, &user.Password)
		if err != nil {
			return nil, fmt.Errorf("users get scan err: %w", err)
		}
		users = append(users, user)
	}

	return users, nil
}
