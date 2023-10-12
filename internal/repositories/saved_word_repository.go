package repositories

import (
	"github.com/frkntplglu/emir-backend/internal/models"
	"gorm.io/gorm"
)

type SavedWordRepository struct {
	db *gorm.DB
}

func NewSavedWordRepository(db *gorm.DB) *SavedWordRepository {
	return &SavedWordRepository{
		db: db,
	}
}

func (r *SavedWordRepository) All(fields models.SavedWord, pagination *models.Pagination) ([]models.SavedWord, error) {
	var words []models.SavedWord
	offset := (pagination.Page - 1) * pagination.Limit
	result := r.db.Table("saved_word").Where(fields).Limit(pagination.Limit).Offset(offset).Find(&words)

	if result.Error != nil {
		return nil, result.Error
	}

	return words, nil
}

func (r *SavedWordRepository) Retrieve(fields models.SavedWord) (models.SavedWord, error) {
	var word models.SavedWord
	result := r.db.Table("saved_word").Where(fields).First(&word)

	if result.Error != nil {
		return models.SavedWord{}, result.Error
	}

	return word, nil
}

func (r *SavedWordRepository) Create(word *models.SavedWord) error {

	result := r.db.Table("saved_word").Create(&word)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *SavedWordRepository) Update(word *models.SavedWord, fields models.SavedWord) error {

	result := r.db.Table("saved_word").Model(&word).Where("id = ?", word.Id).Updates(fields)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *SavedWordRepository) Delete(wordId int) error {

	result := r.db.Table("saved_word").Where("id = ?", wordId).Delete(models.SavedWord{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
