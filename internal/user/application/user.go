// File: internal/user/application/user_service.go

package application

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/pratham704/golang-ddd/internal/user/domain/entity"
	"github.com/pratham704/golang-ddd/internal/user/infrastructure/persistence"
)



func GetUsers(r *http.Request) ([]*entity.ResponseUser, error) {
	// Call repository to get all users
	queryValues := r.URL.Query()
	repo, err := persistence.NewUserRepository()
	if err != nil {
		return nil, err
	}
	users, userErr := repo.GetAllUsers(queryValues)
	if userErr != nil {
		return nil, userErr
	}
	return users, nil
}

// CreateUser creates a new user
func CreateUser(user *entity.User) (*entity.User, error) {
	// Add any validation or business logic here before creating the user

	repo, err := persistence.NewUserRepository()
	if err != nil {
		return nil, err
	}

	user, err2 := repo.CreateUser(user)
	if err2 != nil {
		return nil, err2
	}
	return user, nil
}

// GetUserByID retrieves a user by ID
func GetUserByID(r *http.Request) (*entity.User, error) {
	// Call repository to get user by ID
	repo, err := persistence.NewUserRepository()
	if err != nil {
		return nil, err
	}
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID")
	}
	user, userErr := repo.GetUserByID(uint(id))
	if userErr != nil {
		return nil, userErr
	}
	return user, nil
}
func GetUser(r *http.Request) (*entity.ResponseUser, error) {
	// Call repository to get user by ID
	repo, err := persistence.NewUserRepository()
	if err != nil {
		return nil, err
	}
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID")
	}
	user, userErr := repo.GetUser(uint(id))
	if userErr != nil {
		return nil, userErr
	}
	return user, nil
}

// UpdateUser updates an existing user
func UpdateUser(r *http.Request, user *entity.UpdateUser) (*entity.User, error) {
	// Call repository to update user
	oldUser, err := GetUserByID(r)
	if err != nil {
		return nil, err
	}
	repo, err := persistence.NewUserRepository()
	if err != nil {
		return nil, err
	}

	updateUser, err2 := repo.UpdateUser(oldUser, user)
	if err2 != nil {
		return nil, err2
	}
	return updateUser, nil
}

// DeleteUser deletes a user by ID
func DeleteUser(r *http.Request) error {
	// Call repository to delete user
	user, err := GetUserByID(r)
	if err != nil {
		return err
	}
	repo, err := persistence.NewUserRepository()
	if err != nil {
		return err
	}

	err2 := repo.DeleteUser(user)
	if err2 != nil {
		return err2
	}

	return nil
}

// ChangePassword changes the password of a user
func ChangePassword(r *http.Request, user *entity.UserPasswordChange) error {
	// Call repository to change password
	oldUser, err := GetUserByID(r)
	if err != nil {
		return err
	}
	repo, repoErr := persistence.NewUserRepository()
	if repoErr != nil {
		return repoErr
	}
	userErr := repo.ChangePassword(oldUser, user)
	if userErr != nil {
		return userErr
	}
	return nil
}

// Login authenticates a user
func Login(loginUser *entity.LoginUser) (*entity.LoginUserResponse, error) {
	// Call repository to authenticate user
	repo, err := persistence.NewUserRepository()
	if err != nil {
		return nil,  err
	}
	user, userErr := repo.Login(loginUser)
	if userErr != nil {
		return nil,  userErr
	}
	return user, nil
}