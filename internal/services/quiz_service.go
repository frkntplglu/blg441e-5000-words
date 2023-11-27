package services

import (
	"github.com/frkntplglu/emir-backend/internal/models"
	"github.com/frkntplglu/emir-backend/internal/repositories"
)

type QuizService struct {
	quizRepository repositories.QuizRepository
}

func NewQuizService(quizRepository repositories.QuizRepository) *QuizService {
	return &QuizService{
		quizRepository: quizRepository,
	}
}

func (s *QuizService) GetAllQuizzes(fields models.Quiz, pagination *models.Pagination) ([]models.Quiz, error) {
	return s.quizRepository.All(fields, pagination)
}

func (s *QuizService) GetQuizById(quizId int) (models.Quiz, error) {
	return s.quizRepository.Retrieve(models.Quiz{
		Id: quizId,
	})
}

func (s *QuizService) CreateQuiz(quiz *models.Quiz) error {
	err := s.quizRepository.Create(quiz)
	if err != nil {
		return err
	}

	return nil
}

func (s *QuizService) UpdateQuizById(quiz *models.Quiz, fields models.Quiz) error {
	err := s.quizRepository.Update(quiz, fields)
	if err != nil {
		return err
	}

	return nil
}

func (s *QuizService) DeleteQuizById(quizId int) error {
	err := s.quizRepository.Delete(quizId)
	if err != nil {
		return err
	}

	return nil
}
