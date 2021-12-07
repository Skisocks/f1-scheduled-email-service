package repository

import (
	"database/sql"
	"email-service/models"
	"fmt"
)

type Repository interface {
	All() ([]models.User, error)
}

// repository is a custom type which wraps the sql.DB connection pool REPO
type repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{DB: db}
}

func (repo *repository) All() ([]models.User, error) {
	rows, err := repo.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usrs []models.User

	for rows.Next() {
		var usr models.User

		err := rows.Scan(&usr.UserID, &usr.FirstName, &usr.LastName, &usr.Email)
		if err != nil {
			return nil, err
		}

		usrs = append(usrs, usr)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return usrs, nil
}

func (repo *repository) UserEmails() []models.User {
	usrs, err := repo.All()
	if err != nil {
		return nil
	}

	for _, usr := range usrs {
		fmt.Printf("%s\n", usr.FirstName)
		fmt.Printf("hello")
	}
	return usrs
}
