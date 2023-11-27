package repositories

import (
	"github.com/frkntplglu/emir-backend/internal/models"
	"gorm.io/gorm"
)

type QuizRepository struct {
	db *gorm.DB
}

func NewQuizRepository(db *gorm.DB) *QuizRepository {
	return &QuizRepository{
		db: db,
	}
}

func (r *QuizRepository) All(fields models.Quiz, pagination *models.Pagination) ([]models.Quiz, error) {
	var quizzes []models.Quiz
	offset := (pagination.Page - 1) * pagination.Limit
	result := r.db.Table("quiz").Where(fields).Limit(pagination.Limit).Offset(offset).Find(&quizzes)

	if result.Error != nil {
		return nil, result.Error
	}

	return quizzes, nil
}

func (r *QuizRepository) Retrieve(fields models.Quiz) (models.Quiz, error) {
	var quiz models.Quiz
	result := r.db.Table("quiz").Where(fields).First(&quiz)

	if result.Error != nil {
		return models.Quiz{}, result.Error
	}

	return quiz, nil
}

func (r *QuizRepository) Create(quiz *models.Quiz) error {

	result := r.db.Table("quiz").Create(&quiz)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *QuizRepository) Update(quiz *models.Quiz, fields models.Quiz) error {

	result := r.db.Table("quiz").Model(&quiz).Where("id = ?", quiz.Id).Updates(fields)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *QuizRepository) Delete(quizId int) error {

	result := r.db.Table("quiz").Where("id = ?", quizId).Delete(models.Quiz{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
