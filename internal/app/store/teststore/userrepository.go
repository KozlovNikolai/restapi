package teststore

import (
	"github.com/KozlovNikolai/restapi/internal/app/store"
	"github.com/KozlovNikolai/restapi/model"
)

// UserRepository ...
type UserRepository struct {
	store *Store
	users map[string]*model.User
}

// Create ...
func (ur *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	ur.users[u.Email] = u
	u.ID = len(ur.users)
	return nil
}

// FindByEmail ...
func (ur *UserRepository) FindByEmail(email string) (*model.User, error) {
	u, ok := ur.users[email]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return u, nil
}
