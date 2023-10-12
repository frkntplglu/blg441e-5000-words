package services

import (
	"github.com/frkntplglu/emir-backend/internal/models"
	"github.com/frkntplglu/emir-backend/internal/repositories"
)

type WordService struct {
	wordRepository repositories.WordRepository
}

func NewWordService(wordRepository repositories.WordRepository) *WordService {
	return &WordService{
		wordRepository: wordRepository,
	}
}

func (s *WordService) GetAllWords(fields models.Word, pagination *models.Pagination) ([]models.Word, error) {
	return s.wordRepository.All(fields, pagination)
}

func (s *WordService) GetWordById(wordId int) (models.Word, error) {
	return s.wordRepository.Retrieve(models.Word{
		Id: wordId,
	})
}

func (s *WordService) CreateWord(word *models.Word) error {
	err := s.wordRepository.Create(word)
	if err != nil {
		return err
	}

	return nil
}

func (s *WordService) UpdateWordById(word *models.Word, fields models.Word) error {
	err := s.wordRepository.Update(word, fields)
	if err != nil {
		return err
	}

	return nil
}

func (s *WordService) DeleteWordById(wordId int) error {
	err := s.wordRepository.Delete(wordId)
	if err != nil {
		return err
	}

	return nil
}
