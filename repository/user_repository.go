package repository

import (
	"database/sql"
	"enigma-laundry-apps/model"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Create(payload model.UserCredential) error
	GetUsername(username string) (model.UserCredential, error)
	GetUsernamePassword(username string, password string) (model.UserCredential, error)
	List() ([]model.UserCredential, error)
}

type userRepository struct {
	db *sql.DB
}

// Create implements UserRepository.
func (u *userRepository) Create(payload model.UserCredential) error {
	_, err := u.db.Exec("INSERT INTO user_credential (id, username, password) VALUES ($1, $2, $3)", payload.Id, payload.Username, payload.Password)
	if err != nil {
		return err
	}
	return nil
}

// GetUsername implements UserRepository.
func (u *userRepository) GetUsername(username string) (model.UserCredential, error) {
	var user model.UserCredential
	err := u.db.QueryRow("SELECT id, username, password FROM user_credential WHERE is_active = $1 AND username = $2", true, username).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return model.UserCredential{}, err
	}
	return user, nil
}

// GetUsernamePassword implements UserRepository.
func (u *userRepository) GetUsernamePassword(username string, password string) (model.UserCredential, error) {

	user, err := u.GetUsername(username)
	if err != nil {
		return model.UserCredential{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return model.UserCredential{}, fmt.Errorf("failed to verify password hash : %v", err)
	}

	return user, nil
}

// List implements UserRepository.
func (u *userRepository) List() ([]model.UserCredential, error) {
	var users []model.UserCredential
	rows, err := u.db.Query("SELECT id, username, is_active FROM user_credential")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user model.UserCredential
		err := rows.Scan(&user.Id, &user.Username, &user.IsActive)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
