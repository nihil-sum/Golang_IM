package models

type Session struct {
	UserID uint

	ClientIP string

	ClientPort string

	LoginTime uint64

	HeartBeatTime uint64

	IsOnline bool

	DeviceInfo string
}
