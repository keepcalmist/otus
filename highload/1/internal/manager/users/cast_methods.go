package users

import "github.com/keepcalmist/otus/highload/1/internal/database/models"

func castUsers(users []*models.User) []*User {
	casted := make([]*User, 0, len(users))
	for _, u := range users {
		casted = append(casted, &User{
			ID:         u.ID,
			FirstName:  u.FirstName,
			SecondName: u.SecondName,
			Age:        u.Age,
			Biography:  u.Biography,
			City:       u.City,
		})
	}
	return casted
}
