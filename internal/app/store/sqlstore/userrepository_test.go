package sqlstore_test

import (
	"fmt"
	"testing"

	"github.com/KozlovNikolai/restapi/internal/app/store"
	"github.com/KozlovNikolai/restapi/internal/app/store/sqlstore"
	"github.com/KozlovNikolai/restapi/model"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL) //открываем БД
	defer teardown("users")                         //чистим таблицу при выходе, и закрываем БД

	s := sqlstore.New(db) //создаем хранилище, подключенное к БД
	email := "user@example.org"
	_, err := s.User().FindByEmail(email)                      //ищем несуществющего пользователя
	assert.EqualError(t, err, store.ErrRecordNotFound.Error()) //проверяем, что такого пользователя нет в базе

	u := model.TestUser(t) //создаем тестового пользователя
	u.Email = email
	err = s.User().Create(u)             //создаем запись пользователя в БД
	assert.NoError(t, err)               //проверяем отсутствие ошибок при создании записи
	u, err = s.User().FindByEmail(email) //проверяем поиск пользователя по email
	assert.NoError(t, err)               //проверяем отсутсвие ошибок поиска
	assert.NotNil(t, u)                  //проверяем, что найденый пользователь не пустой
}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")
	s := sqlstore.New(db)
	id := 0
	_, err := s.User().Find(id)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())
	fmt.Println("test \"non-existent user\" is stoped")

	u1 := model.TestUser(t)
	err = s.User().Create(u1)
	assert.NoError(t, err)
	u1, err = s.User().Find(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u1)
}
