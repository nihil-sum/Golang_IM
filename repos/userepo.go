package repos

import (
	"context"

	"github.com/nihil_sum/Golang_IM/models"
	"github.com/nihil_sum/Golang_IM/seed"
	"gorm.io/gorm"
)

var db *gorm.DB = seed.Init()

type UsersRepo interface {
	MatchByUsername(username string, password string) (models.UserBasic, error) //Find the user in database
	CreateUser(user models.UserBasic) error                                     //Register
	DeleteUser(username string) error                                           //Delete the account
}
type urepo struct {
}

func (r *urepo) MatchByUsername(username string, password string) (
	models.UserBasic,
	error,
) {
	ctx := context.Background()
	user, err := gorm.G[models.UserBasic](db).First(ctx)
	if err != nil {
		panic("Searching error in database")
	}
	return user, nil
}
func CreateUser(user models.UserBasic) error {

	ctx := context.Background()
	err := gorm.G[models.UserBasic](db).Create(ctx, &user)
	if err != nil {
		return err
	}

	return nil
}
