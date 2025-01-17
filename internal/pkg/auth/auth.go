package auth

import (
	"authService/internal/api/controllers/authController/dto"
	"authService/internal/models"
	"authService/internal/pkg/crypto"
	"authService/internal/pkg/token"
	"authService/internal/storage/postgres"
	"errors"
)

type Service struct {
	Storage      *postgres.Postgres
	TokenService *token.Service
}

func NewService(storage *postgres.Postgres, tokenService *token.Service) *Service {
	return &Service{
		Storage:      storage,
		TokenService: tokenService,
	}
}

func (s *Service) Login(req dto.LoginReq, responseChan chan *dto.ReturnToken) {

	username := req.Username
	password := req.Password

	user, err := s.Storage.GetUserByUsername(username)
	if err != nil {
		responseChan <- &dto.ReturnToken{
			Token: nil,
			Error: errors.New("user not found"),
		}
		return
	}

	if crypto.CheckPasswordHash(password, user.Password) {
		accessToken, err := s.TokenService.GenerateAccessToken(user)
		if err != nil {
			responseChan <- &dto.ReturnToken{
				Token: nil,
				Error: errors.New("could not generate access token"),
			}
			return
		}

		refreshToken, err := s.TokenService.GenerateRefreshToken(user)
		if err != nil {
			responseChan <- &dto.ReturnToken{
				Token: nil,
				Error: errors.New("could not generate refresh token"),
			}
			return
		}

		tokenObj := &dto.Token{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}

		responseChan <- &dto.ReturnToken{
			Token: tokenObj,
			Error: nil,
		}

	} else {
		responseChan <- &dto.ReturnToken{
			Token: nil,
			Error: errors.New("invalid password"),
		}
		return
	}
}

func (s *Service) Register(req dto.RegisterReq, responseChan chan *dto.ReturnConfirmation) {

	username := req.Username
	password := req.Password
	email := req.Email
	firstname := req.Firstname
	lastname := req.Lastname

	hashedPassword, err := crypto.HashPassword(password)
	if err != nil {
		responseChan <- &dto.ReturnConfirmation{
			Confirmation: false,
			Error:        errors.New("could not hash password"),
		}
		return
	}

	user := &models.User{
		Username:  username,
		Password:  hashedPassword,
		Email:     email,
		Firstname: firstname,
		Lastname:  lastname,
	}

	_, err = s.Storage.GetUserByUsername(username)
	if err == nil {
		responseChan <- &dto.ReturnConfirmation{
			Confirmation: false,
			Error:        errors.New("user already exists"),
		}
		return
	}

	err = s.Storage.CreateUser(user)
	if err != nil {
		responseChan <- &dto.ReturnConfirmation{
			Confirmation: false,
			Error:        errors.New("could not create user"),
		}
		return
	}

	responseChan <- &dto.ReturnConfirmation{
		Confirmation: true,
		Error:        nil,
	}

}
