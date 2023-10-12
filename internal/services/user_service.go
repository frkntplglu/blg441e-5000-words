package services

import (
	"github.com/frkntplglu/emir-backend/internal/models"
	"github.com/frkntplglu/emir-backend/internal/repositories"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) GetAllUsers(fields models.User, pagination *models.Pagination) ([]models.User, error) {
	return s.userRepository.All(fields, pagination)
}

func (s *UserService) GetUserByParams(fields models.User) (models.User, error) {
	return s.userRepository.Retrieve(fields)
}

func (s *UserService) GetUserById(userId int) (models.User, error) {
	return s.userRepository.Retrieve(models.User{
		Id: userId,
	})
}

func (s *UserService) CreateUser(user *models.User) error {
	err := s.userRepository.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) UpdateUserById(user *models.User, fields models.User) error {
	err := s.userRepository.Update(user, fields)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) DeleteUserById(userId int) error {
	err := s.userRepository.Delete(userId)
	if err != nil {
		return err
	}

	return nil
}
