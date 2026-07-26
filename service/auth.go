package service

import "github.com/nihil_sum/Golang_IM/models"

type LoginService struct{}

// @TODO
func (s *LoginService) Login(username string, password string) (models.UserBasic, error) {
	var user models.UserBasic
	return user, nil
}
