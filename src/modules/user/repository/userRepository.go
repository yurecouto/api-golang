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
		`INSERT INTO "users"("name", "email", "password", "created_at")
    values($1, $2, $3, $4)`,
	)

	if erro != nil {
		return false, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(
		user.Name,
		user.Email,
		user.Password,
		user.CreatedAt,
	)
	if erro != nil {
		return false, erro
	}

	return result != nil, nil
}

func (repo Users) FindAllUsers() ([]models.User, error) {
	lines, erro := repo.db.Query(
		"SELECT id, name, email, created_at FROM users",
	)
	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}

	return users, nil
}

func (repo Users) FindById(userID uint64) (models.User, error) {
	lines, erro := repo.db.Query(
		"SELECT id, name, email, created_at FROM users WHERE id = $1",
		userID,
	)
	if erro != nil {
		return models.User{}, erro
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return models.User{}, erro
		}
	}

	return user, nil
}
