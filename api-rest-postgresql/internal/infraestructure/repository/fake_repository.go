package repository

import "api-rest-postgresql/internal/domain/user"

type FakeUserRepository struct{}

func (r *FakeUserRepository) FindAll() ([]user.User, error) {
	return []user.User{
		{ID: 1, Name: "John Doe", Email: "john.doe@example.com"},
		{ID: 2, Name: "Jane Doe", Email: "jane.doe@example.com"},
	}, nil
}

func (r *FakeUserRepository) FindByID(id int) (*user.User, error) {
	return nil, nil
}

func (r *FakeUserRepository) Create(user *user.User) error {
	return nil
}

func (r *FakeUserRepository) Update(user *user.User) error {
	return nil
}

func (r *FakeUserRepository) Delete(id int) error {
	return nil
}
