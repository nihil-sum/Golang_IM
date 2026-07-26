package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nihil_sum/Golang_IM/router"
	"github.com/nihil_sum/Golang_IM/seed"
)

func main() {
	seed.Init()
	r := gin.Default()
	router.SetupRouter(r)

}
