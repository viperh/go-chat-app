package users

import (
	"authService/internal/api/controllers/usersController/dto"
	"authService/internal/models"
	"authService/internal/storage/postgres"
)

type Service struct {
	Storage *postgres.Postgres
}

func NewService(storage *postgres.Postgres) *Service {
	return &Service{
		Storage: storage,
	}
}

func (s *Service) GetUserByID(id int, responseChan chan *dto.ReturnUser) {
	user, err := s.Storage.GetUserById(id)

	if err != nil {
		responseChan <- &dto.ReturnUser{
			User:  nil,
			Error: err,
		}
		return
	}

	responseChan <- &dto.ReturnUser{
		User:  user,
		Error: nil,
	}

}

func (s *Service) GetUserByEmail(email string, responseChan chan *dto.ReturnUser) {
	user, err := s.Storage.GetUserByEmail(email)
	if err != nil {
		responseChan <- &dto.ReturnUser{
			User:  nil,
			Error: err,
		}
		return
	}

	responseChan <- &dto.ReturnUser{
		User:  user,
		Error: nil,
	}
}

func (s *Service) GetUserByUsername(username string, responseChan chan *dto.ReturnUser) {
	user, err := s.Storage.GetUserByUsername(username)
	if err != nil {
		responseChan <- &dto.ReturnUser{
			User:  nil,
			Error: err,
		}
		return
	}

	responseChan <- &dto.ReturnUser{
		User:  user,
		Error: nil,
	}
}

func (s *Service) CreateUser(user *models.User, responseChan chan *dto.ReturnConfirmation) {
	err := s.Storage.CreateUser(user)
	if err != nil {
		responseChan <- &dto.ReturnConfirmation{
			Confirmation: false,
			Error:        err,
		}
		return
	}

	responseChan <- &dto.ReturnConfirmation{
		Confirmation: true,
		Error:        nil,
	}
}

func (s *Service) UpdateUser(user *models.User, responseChan chan *dto.ReturnConfirmation) {
	err := s.Storage.UpdateUser(user)
	if err != nil {
		responseChan <- &dto.ReturnConfirmation{
			Confirmation: false,
			Error:        err,
		}
		return
	}

	responseChan <- &dto.ReturnConfirmation{
		Confirmation: true,
		Error:        nil,
	}

}

func (s *Service) DeleteUser(id int, responseChan chan *dto.ReturnConfirmation) {
	err := s.Storage.DeleteUser(id)

	if err != nil {
		responseChan <- &dto.ReturnConfirmation{
			Confirmation: false,
			Error:        err,
		}
		return
	}

	responseChan <- &dto.ReturnConfirmation{
		Confirmation: true,
		Error:        nil,
	}

}
