package models

type Message struct {
	Buf []byte //The message data buffer
	Dsn string //The message destination ip
	To  bool   //Broadcast if true

}
