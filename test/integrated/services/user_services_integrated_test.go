package services_test

import (
	"database/sql"
	"testing"

	"github.com/okyws/dashboard-backend/adapter/repository"
	"github.com/okyws/dashboard-backend/domain"
	"github.com/okyws/dashboard-backend/services"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	assert.NoError(t, err)
	assert.NotNil(t, db)

	// check error close handling
	defer func() {
		if err := db.Close(); err != nil {
			assert.NoError(t, err)
		}
	}()

	gormDB, err := gorm.Open(sqlite.New(sqlite.Config{Conn: db}), &gorm.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, gormDB)

	userRepository := repository.NewUserRepositoryAdapter(gormDB)
	userService := services.NewUserService(userRepository)

	t.Run("Empty password", func(t *testing.T) {
		user := &domain.User{Username: "testuser", Password: ""}
		_, err := userService.CreateUser(user)

		assert.Error(t, err)
		assert.Equal(t, "password cannot be empty", err.Error())
	})

	t.Run("Existing username", func(t *testing.T) {
		gormDB.Exec("DROP TABLE users")
		err := gormDB.AutoMigrate(&domain.User{})
		assert.NoError(t, err)

		existingUser := &domain.User{Username: "testuser", Password: "testpassword"}
		gormDB.Create(existingUser)

		user := &domain.User{Username: "testuser", Password: "testpassword"}
		_, err = userService.CreateUser(user)

		assert.Error(t, err)
		assert.Equal(t, "username already exists", err.Error())
	})

	t.Run("Successful user creation", func(t *testing.T) {
		gormDB.Exec("DROP TABLE users")
		err := gormDB.AutoMigrate(&domain.User{})
		assert.NoError(t, err)

		user := &domain.User{Username: "testuser", Password: "testpassword"}
		createdUser, err := userService.CreateUser(user)
		assert.NoError(t, err)
		assert.NotNil(t, createdUser)
		assert.Equal(t, "", createdUser.Password)
	})

	t.Run("Database error", func(t *testing.T) {
		gormDB.Exec("DROP TABLE users")

		user := &domain.User{Username: "testuser", Password: "testpassword"}
		createdUser, err := userService.CreateUser(user)
		assert.Error(t, err)
		assert.Nil(t, createdUser)
		assert.Contains(t, err.Error(), "no such table: users")
	})
}
