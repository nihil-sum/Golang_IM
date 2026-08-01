package websocket

import (
	"github.com/nihil_sum/Golang_IM/models"
	"golang.org/x/net/websocket"
)

type Client struct {
	Send chan []byte
	User *models.UserBasic
	Conn *websocket.Conn
}

func SendMessage(models.Message) {
	
}
