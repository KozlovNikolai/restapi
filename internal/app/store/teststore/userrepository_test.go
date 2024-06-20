package teststore_test

import (
	"fmt"
	"testing"

	"github.com/KozlovNikolai/restapi/internal/app/store"
	"github.com/KozlovNikolai/restapi/internal/app/store/teststore"
	"github.com/KozlovNikolai/restapi/model"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {

	s := teststore.New()
	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email
	err = s.User().Create(u)
	assert.NoError(t, err)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {
	s := teststore.New()

	u := model.TestUser(t)
	err := s.User().Create(u)
	assert.NoError(t, err)

	_, err = s.User().Find(u.ID + 1)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())
	fmt.Println("test \"non-existent user\" is stoped")

	u, err = s.User().Find(u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u)
	fmt.Println("test \"existent user\" is stoped")
}
