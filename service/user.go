package service

import (
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func (us *UserService) Register(username, password, nickname string) bool {
	return true
}
