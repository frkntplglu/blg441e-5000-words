package services

import (
	"github.com/frkntplglu/emir-backend/internal/models"
	"github.com/frkntplglu/emir-backend/internal/repositories"
)

type SavedWordService struct {
	savedWordRepository repositories.SavedWordRepository
}

func NewSavedWordService(savedWordRepository repositories.SavedWordRepository) *SavedWordService {
	return &SavedWordService{
		savedWordRepository: savedWordRepository,
	}
}

func (s *SavedWordService) GetAllWords(fields models.SavedWord, pagination *models.Pagination) ([]models.SavedWord, error) {
	return s.savedWordRepository.All(fields, pagination)
}

func (s *SavedWordService) CreateSavedWord(savedWord *models.SavedWord) error {
	err := s.savedWordRepository.Create(savedWord)
	if err != nil {
		return err
	}

	return nil
}

/*
func (s *SavedWordService) GetWordById(wordId int) (models.Word, error) {
	return s.wordRepository.Retrieve(models.Word{
		Id: wordId,
	})
}

func (s *SavedWordService) CreateWord(word *models.Word) error {
	err := s.wordRepository.Create(word)
	if err != nil {
		return err
	}

	return nil
}

func (s *SavedWordService) UpdateWordById(word *models.Word, fields models.Word) error {
	err := s.wordRepository.Update(word, fields)
	if err != nil {
		return err
	}

	return nil
}

func (s *SavedWordService) DeleteWordById(wordId int) error {
	err := s.wordRepository.Delete(wordId)
	if err != nil {
		return err
	}

	return nil
}
*/
