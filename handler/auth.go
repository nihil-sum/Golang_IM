package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nihil_sum/Golang_IM/repos"
	"github.com/nihil_sum/Golang_IM/service"
)
type Claims struct {
	UserID uint `json:"user_id"`
	UserName string `json:"username"`
	jwt.RegisteredClaims
}
type LoginRequest struct {
	username string `json:"username"`
	password string `json:"password"`
}
type LoginHandler struct {
	s *service.LoginService
}
func GenerateToken(userID uint,userName string) (string,error){
	claims := Claims{
		UserID: userID,
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiredAt,
		},
	}
}
func (h *LoginHandler) LoginHandler(c *gin.Context) {
	var req LoginRequest
	c.ShouldBindJSON(&req)
	user, err := h.s.Login(
		req.username,
		req.password,
	)
	if err != nil {
		panic("internal service error")
	}
	repos.MatchByUsername(user.Name,user.Password)
	c.JSON(
		http.StatusOK,
		gin.H{
			token: ,
		},
	)
}
