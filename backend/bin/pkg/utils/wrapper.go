package utils

import (
	"login-api-jwt/bin/modules/user/models"

	"gorm.io/gorm"
)

type Result struct {
	Data  interface{}
	DB    *gorm.DB
	Error error
}

type FindPasswordResult struct {
	Data     models.User
	Password string
	DB       *gorm.DB
	Error    error
}
