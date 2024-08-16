package seeders

import (
	entity "github.com/123-zuleyha/go_backend_project/entity"
)

type UserSeeder struct{}

func (u *UserSeeder) Run() []entity.User {
	seed := []entity.User{
		{
			Username:   "admin",
			UserTypeID: 1,
			Email:      "admin@admin.com",
			FirstName:  "admin",
			LastName:   "admin",
			Password:   "12345",
		},
	}
	return seed
}
