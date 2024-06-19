package teststore

import (
	"github.com/KozlovNikolai/restapi/internal/app/store"
	"github.com/KozlovNikolai/restapi/model"
)

// UserRepository ...
type UserRepository struct {
	store *Store
	users map[int]*model.User
}

// Create ...
func (ur *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}
	u.ID = len(ur.users) + 1
	ur.users[u.ID] = u
	return nil
}

// FindByEmail ...
func (ur *UserRepository) FindByEmail(email string) (*model.User, error) {
	for _, u := range ur.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, store.ErrRecordNotFound
}

// Find ...
func (ur *UserRepository) Find(id int) (*model.User, error) {
	u, ok := ur.users[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return u, nil
}
