package userrepository

import (
	"api-golang/src/models"
	"api-golang/src/utils"
	"database/sql"
	"fmt"
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

	if user.CreatedAt.String() == "0001-01-01 00:00:00 +0000 UTC" {
		return user, fmt.Errorf("No user was Found.")
	}

	return user, nil
}

func (repo Users) FindByIdAndUpdate(userID uint64, data models.User) (
	models.User, error) {

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

	if user.CreatedAt.String() == "0001-01-01 00:00:00 +0000 UTC" {
		return user, fmt.Errorf("No user was Found.")
	}

	statement, erro := repo.db.Prepare(
		"UPDATE users SET name = $1, password = $2, email = $3 WHERE id = $4",
	)
	if erro != nil {
		return data, erro
	}
	defer statement.Close()

	var updatedUser models.User

	updatedUser.ID = user.ID
	updatedUser.CreatedAt = user.CreatedAt

	if data.Name == "" {
		updatedUser.Name = user.Name
	} else {
		updatedUser.Name = data.Name
	}

	if data.Password == "" {
		updatedUser.Password = user.Password
	} else {
		passwordHash, erro := utils.HashPassword(data.Password)
		if erro != nil {
			return data, erro
		}

		updatedUser.Password = string(passwordHash)
	}

	if data.Email == "" {
		updatedUser.Email = user.Email
	} else {
		updatedUser.Email = data.Email
	}

	if _, erro := statement.Exec(
		updatedUser.Name,
		updatedUser.Password,
		updatedUser.Email,
		userID); erro != nil {
		return data, erro
	}

	return updatedUser, nil
}

func (repo Users) Delete(userID uint64) error {
	statement, erro := repo.db.Prepare("DELETE FROM users WHERE id = $1")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(userID); erro != nil {
		return erro
	}

	return nil
}

func (repo Users) FindByEmail(email string) (models.User, error) {
	lines, erro := repo.db.Query(
		"SELECT id, name, email, password, created_at FROM users WHERE email = $1",
		email,
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
			&user.Password,
			&user.CreatedAt,
		); erro != nil {
			return models.User{}, erro
		}
	}

	if user.CreatedAt.String() == "0001-01-01 00:00:00 +0000 UTC" {
		return user, fmt.Errorf("No user was Found.")
	}

	return user, nil
}

func (repo Users) SaveRefreshToken(token string, id int32) (bool, error) {
	statement, erro := repo.db.Prepare(
		`INSERT INTO "user_tokens"("token", "user_id") values($1, $2)`,
	)

	if erro != nil {
		return false, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(
		token,
		id,
	)
	if erro != nil {
		return false, erro
	}

	return result != nil, nil
}

func (repo Users) FindRefreshToken(
	token string,
) (models.UserToken, error) {

	lines, erro := repo.db.Query(
		"SELECT id, token, user_id FROM user_tokens WHERE token = $1",
		token,
	)
	if erro != nil {
		return models.UserToken{}, erro
	}
	defer lines.Close()

	var userToken models.UserToken

	if lines.Next() {
		if erro = lines.Scan(
			&userToken.ID,
			&userToken.Token,
			&userToken.UserId,
		); erro != nil {
			return models.UserToken{}, fmt.Errorf("No refresh token was Found.")
		}
	}

	return userToken, nil
}
