package repositories

import (
	"database/sql"
	"email-service/models"
	"fmt"
	"go.uber.org/zap"
)

type UserEmailGetter interface {
	GetUserEmails() []string
}

// repository is a custom type which wraps the sql.DB connection pool REPO
type repository struct {
	logger *zap.Logger
	DB     *sql.DB
}

func NewRepository(logger *zap.Logger, db *sql.DB) *repository {
	return &repository{
		logger: logger,
		DB:     db,
	}
}

// GetAll returns all the information regarding all the current users of the service
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

// GetUserEmails returns the email addresses of all the current users of the service
func (repo *repository) GetUserEmails() []string {
	users, err := repo.GetAll()
	if err != nil {
		repo.logger.Error(fmt.Sprintf("failed to query database: %s", err.Error()))
	}

	var userEmails []string
	for _, user := range users {
		userEmails = append(userEmails, user.Email)
	}
	return userEmails
}
