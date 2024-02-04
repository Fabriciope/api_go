package repositories

import (
	"database/sql"
	"testing"

	"github.com/Fabriciope/my-api/configs"
	"github.com/Fabriciope/my-api/internal/infra/database"
	"github.com/Fabriciope/my-api/internal/models"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

var conn *sql.DB

func init() {
	configs.Cfg = &configs.Config{
		DBDriver: "mysql",
		DBHost: "localhost",
		DBPort: 7000,
		DBUser: "root",
		DBName: "api_golang",
		DBPassword: "password",
	}

	conn, _ = database.Connect()
}


func TestInsertNewUser(t *testing.T) {
	repository, _ := NewRepository(conn)

	user, _ := models.NewUser(gofakeit.Name(), gofakeit.Email(), "password")
	err := repository.User.Create(user)

	assert.Nil(t, err)

	_, err = repository.User.FindOneWhere("id", user.ID)
	
	assert.Nil(t, err)
}

func TestUpdateAUser(t *testing.T) {
	repository, _ := NewRepository(conn)
	
	user, _ := models.NewUser(gofakeit.Name(), gofakeit.Email(), "password")
	repository.User.Create(user)

	userFound, _ := repository.User.FindOneWhere("id", user.ID)
	u := userFound.(*models.User)
	u.Name = "New name: " + gofakeit.Name()
	
	err := repository.User.Update(u)
	
	assert.Nil(t, err)
	
	userUpdated, err := repository.User.FindOneWhere("name", u.Name)
	
	assert.NotNil(t, userUpdated)
	assert.Nil(t, err)
}

func TestDeleteAUser(t *testing.T) {
	repository, _ := NewRepository(conn)
	
	user, _ := models.NewUser(gofakeit.Name(), gofakeit.Email(), "password")
	repository.User.Create(user)
	
	err := repository.User.Delete(user.ID)
	
	assert.Nil(t, err)
	
	userDeleted, err := repository.User.FindOneWhere("id", user.ID)
	
	assert.NotNil(t, err)
	assert.Nil(t, userDeleted)
}

func TestFindAllUsersWithPagination(t *testing.T) {
	repository, _ := NewRepository(conn)

	for i := 0; i <= 10; i++ {
		user, _ := models.NewUser(gofakeit.Name(), gofakeit.Email(), "password")
		repository.User.Create(user)
	}

	page, limit := 3, 2
	usersFound, err := repository.User.FindAllWithPagination(page, limit, "desc")
	
	assert.Nil(t, err)
	assert.Len(t, usersFound, limit)
}

func TestFindOneUser(t *testing.T) {
	repository, _ := NewRepository(conn)

	user, _ := models.NewUser(gofakeit.Name(), gofakeit.Email(), "password")
	repository.User.Create(user)

	userFound, err := repository.User.FindOneWhere("email", user.Email)

	assert.Nil(t, err)
	assert.NotNil(t, userFound)
	u := userFound.(*models.User)

	assert.Equal(t, user.ID, u.ID)
	assert.Equal(t, user.Name, u.Name)
	assert.Equal(t, user.Email, u.Email)
	assert.Equal(t, user.Password, u.Password)
	assert.Equal(t, user.CreatedAt, u.CreatedAt)
}

func TestFindOneUserWhenDataIsWrong(t *testing.T) {
	repository, _ := NewRepository(conn)

	userFound, err := repository.User.FindOneWhere("email", "wrongEmail@gmail.com")

	assert.Nil(t, userFound)
	assert.NotNil(t, err)
}

