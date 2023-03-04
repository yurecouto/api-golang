package userrepository

import (
	"api-golang/src/models"
	"database/sql"
)

type Users struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (repo Users) Create(user models.User) (bool, error) {
	statement, erro := repo.db.Prepare(
		`INSERT INTO "users"("name", "email", "password") values($1, $2, $3) RETURNING id`,
	)

	if erro != nil {
		return false, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(user.Name, user.Email, user.Password)
	if erro != nil {
		return false, erro
	}

	return result != nil, nil
}
