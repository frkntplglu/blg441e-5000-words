package services

import (
	"github.com/frkntplglu/emir-backend/internal/models"
	"github.com/frkntplglu/emir-backend/internal/repositories"
)

type QuestionService struct {
	questionRepository repositories.QuestionRepository
}

func NewQuestionService(questionRepository repositories.QuestionRepository) *QuestionService {
	return &QuestionService{
		questionRepository: questionRepository,
	}
}

func (s *QuestionService) GetAllQuestions(fields models.Question, pagination *models.Pagination) ([]models.Question, error) {
	return s.questionRepository.All(fields, pagination)
}

func (s *QuestionService) GetQuestionById(questionId int) (models.Question, error) {
	return s.questionRepository.Retrieve(models.Question{
		Id: questionId,
	})
}

func (s *QuestionService) GetQuestionsByQuizId(quizId int, pagination *models.Pagination) ([]models.Question, error) {
	return s.questionRepository.All(models.Question{
		QuizId: quizId,
	}, pagination)
}

func (s *QuestionService) CreateQuestion(question *models.Question) error {
	err := s.questionRepository.Create(question)
	if err != nil {
		return err
	}

	return nil
}

func (s *QuestionService) UpdateQuestionById(question *models.Question, fields models.Question) error {
	err := s.questionRepository.Update(question, fields)
	if err != nil {
		return err
	}

	return nil
}

func (s *QuestionService) DeleteQuestionById(questionId int) error {
	err := s.questionRepository.Delete(questionId)
	if err != nil {
		return err
	}

	return nil
}
