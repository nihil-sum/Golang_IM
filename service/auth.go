package service

import (
	"github.com/nihil_sum/Golang_IM/models"
	"github.com/nihil_sum/Golang_IM/repos"
)

type LoginService struct {
	R repos.UsersRepo
}
type RegService struct {
	R repos.UsersRepo
}

// @TODO:
func (s *LoginService) Login(
	username string,
	password string,
) (models.UserBasic, error) {

	user, err := s.R.MatchByUsername(
		username,
		password,
	)
	if err != nil {
		return models.UserBasic{}, err
	}

	return user, nil
}
func (s *RegService) Register(
	username string,
	password string,
	email string,
) error {
	err := s.R.CreateUser(models.UserBasic{
		Name:     username,
		Password: password,
		Email:    email,
	})
	if err != nil {
		panic("Failed to create user")
	}
	return nil
}
