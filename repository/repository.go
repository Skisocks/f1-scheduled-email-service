package repository

import (
	"database/sql"
	"email-service/models"
)

type Repository interface {
	GetAll() ([]models.User, error)
	GetUserEmails() []string
}

// repository is a custom type which wraps the sql.DB connection pool REPO
type repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{DB: db}
}

func (repo *repository) GetAll() ([]models.User, error) {
	rows, err := repo.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		err := rows.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *repository) GetUserEmails() []string {
	users, err := repo.GetAll()
	if err != nil {
		return nil
	}

	var userEmails []string
	for _, user := range users {
		userEmails = append(userEmails, user.Email)
	}
	return userEmails
}
