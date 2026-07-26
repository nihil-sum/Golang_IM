package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Name string
	Password string
	Phone string
	Email string
	Identity string

	ClientIP string
	ClientPort string
	LoginTime uint64
	HeartBeatTime uint64
	LogOutTime uint64
	IsLogOut bool
	DeviceInfo string
	
}

func (table *UserBasic) UserBasic() string {
	return "user_basic"
}