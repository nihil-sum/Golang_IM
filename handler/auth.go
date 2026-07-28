package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nihil_sum/Golang_IM/confs"
	"github.com/nihil_sum/Golang_IM/service"
)

type Claims struct {
	UserID   uint   `json:"user_id"`
	UserName string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginHandler struct {
	s *service.LoginService
}

// The
func GenerateToken(userID uint, userName string, role string) (string, error) {
	//Use jwt.New() method to quickly generate a new token.
	//The argument of the jwt.New() is the signing method

	claim := Claims{
		UserID:   userID,
		UserName: userName,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(
					24 * time.Hour,
				),
			),
			IssuedAt: jwt.NewNumericDate(
				time.Now(),
			),
		},
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodEdDSA,
		claim,
	)
	tokenString, err := token.SignedString(confs.PrivateKey)

	if err != nil {
		panic("Failed to sign")
	}
	return tokenString, nil
}
func (h *LoginHandler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "invalid request",
		})
	}
	user, err := h.s.Login(
		req.Username,
		req.Password,
	)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Login Failed",
		})
	}
	token, err := GenerateToken(user.ID, user.Name, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to generate the token",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		token: token,
	})
}
