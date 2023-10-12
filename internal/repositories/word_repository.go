package repositories

import (
	"github.com/frkntplglu/emir-backend/internal/models"
	"gorm.io/gorm"
)

type WordRepository struct {
	db *gorm.DB
}

func NewWordRepository(db *gorm.DB) *WordRepository {
	return &WordRepository{
		db: db,
	}
}

func (r *WordRepository) All(fields models.Word, pagination *models.Pagination) ([]models.Word, error) {
	var words []models.Word
	offset := (pagination.Page - 1) * pagination.Limit
	result := r.db.Table("word").Where(fields).Limit(pagination.Limit).Offset(offset).Find(&words)

	if result.Error != nil {
		return nil, result.Error
	}

	return words, nil
}

func (r *WordRepository) Retrieve(fields models.Word) (models.Word, error) {
	var word models.Word
	result := r.db.Table("word").Where(fields).First(&word)

	if result.Error != nil {
		return models.Word{}, result.Error
	}

	return word, nil
}

func (r *WordRepository) Create(word *models.Word) error {

	result := r.db.Table("word").Create(&word)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *WordRepository) Update(word *models.Word, fields models.Word) error {

	result := r.db.Table("words").Model(&word).Where("id = ?", word.Id).Updates(fields)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *WordRepository) Delete(wordId int) error {

	result := r.db.Table("words").Where("id = ?", wordId).Delete(models.Word{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
