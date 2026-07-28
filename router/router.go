package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nihil_sum/Golang_IM/handler"
)

func SetupRouter(r *gin.Engine) {

	loginHandler := handler.LoginHandler{}
	regHandler := handler.RegHandler{}
	api := r.Group("/api")
	auth := api.Group("/auth")
	{
		auth.POST("/login", loginHandler.Login)
		auth.POST("/register", regHandler.Register)
	}

}
