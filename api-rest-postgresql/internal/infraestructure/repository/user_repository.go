package repository

import (
	"api-rest-postgresql/internal/domain/user"
	"errors"
	"sync"
)

type UserRepository struct {
	users  []user.User
	nextID int
	mutex  sync.RWMutex
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users:  make([]user.User, 0),
		nextID: 1,
	}
}

func (r *UserRepository) FindAll() ([]user.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.users, nil
}

func (r *UserRepository) FindByID(id int) (*user.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, u := range r.users {
		if u.ID == id {
			return &u, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *UserRepository) Create(u *user.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	u.ID = r.nextID
	r.nextID++
	r.users = append(r.users, *u)
	return nil
}

func (r *UserRepository) Update(u *user.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i, existing := range r.users {
		if existing.ID == u.ID {
			r.users[i] = *u
			return nil
		}
	}
	return errors.New("user not found")
}

func (r *UserRepository) Delete(id int) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i, u := range r.users {
		if u.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}
