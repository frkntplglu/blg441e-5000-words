package repositories

import (
	"github.com/frkntplglu/emir-backend/internal/models"
	"gorm.io/gorm"
)

type QuestionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) *QuestionRepository {
	return &QuestionRepository{
		db: db,
	}
}

func (r *QuestionRepository) All(fields models.Question, pagination *models.Pagination) ([]models.Question, error) {
	var questions []models.Question
	offset := (pagination.Page - 1) * pagination.Limit
	result := r.db.Table("question").Where(fields).Limit(pagination.Limit).Offset(offset).Find(&questions)

	if result.Error != nil {
		return nil, result.Error
	}

	return questions, nil
}

func (r *QuestionRepository) Retrieve(fields models.Question) (models.Question, error) {
	var question models.Question
	result := r.db.Table("question").Where(fields).First(&question)

	if result.Error != nil {
		return models.Question{}, result.Error
	}

	return question, nil
}

func (r *QuestionRepository) Create(question *models.Question) error {

	result := r.db.Table("question").Create(&question)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *QuestionRepository) Update(question *models.Question, fields models.Question) error {

	result := r.db.Table("question").Model(&question).Where("id = ?", question.Id).Updates(fields)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *QuestionRepository) Delete(questionId int) error {

	result := r.db.Table("question").Where("id = ?", questionId).Delete(models.Question{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
