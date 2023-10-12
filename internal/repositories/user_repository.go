package repositories

import (
	"github.com/frkntplglu/emir-backend/internal/models"
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

func (r *UserRepository) All(fields models.User, pagination *models.Pagination) ([]models.User, error) {
	var users []models.User
	offset := (pagination.Page - 1) * pagination.Limit
	result := r.db.Table("user").Select("id, firstname, lastname, email, address, last_login, created_at").Where(fields).Limit(pagination.Limit).Offset(offset).Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (r *UserRepository) Retrieve(fields models.User) (models.User, error) {
	var user models.User
	result := r.db.Table("user").Select("id, firstname, lastname, email, address, last_login, created_at").Where(fields).First(&user)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func (r *UserRepository) Create(user *models.User) error {

	result := r.db.Table("user").Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserRepository) Update(user *models.User, fields models.User) error {

	result := r.db.Table("user").Model(&user).Where("id = ?", user.Id).Updates(fields)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserRepository) Delete(userId int) error {

	result := r.db.Table("user").Where("id = ?", userId).Delete(models.User{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
