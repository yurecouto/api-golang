package userrepository

import (
	"api-golang/src/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user *models.User) error {
	result := r.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) Update(id uint64, updatedUser *models.User) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	if err := r.db.Model(&user).Updates(updatedUser).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Delete(id uint64) error {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindById(id uint64) (*models.User, error) {
	var user models.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) SaveRefreshToken(id uint64, token string) error {
	var userToken = models.UserToken{
		UserId: id,
		Token:  token,
	}
	if err := r.db.Create(&userToken).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindRefreshToken(token string) (*models.UserToken, error) {
	var userToken models.UserToken
	if err := r.db.Where("token = ?", token).First(&userToken).Error; err != nil {
		return nil, err
	}
	return &userToken, nil
}

func (r *UserRepository) DeleteRefreshToken(token string) error {
	if err := r.db.Where("token = ?", token).Delete(&models.UserToken{}).Error; err != nil {
		return err
	}
	return nil
}
