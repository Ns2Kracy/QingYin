package service

import "gorm.io/gorm"

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}
