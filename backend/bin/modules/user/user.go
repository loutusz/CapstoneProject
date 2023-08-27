package user

import (
	"login-api-jwt/bin/modules/user/models"
	"login-api-jwt/bin/pkg/utils"

	"github.com/gin-gonic/gin"
)

type UsecaseQuery interface {
	GetByID(ctx *gin.Context)
	GetAccess(ctx *gin.Context)
	GetByName(ctx *gin.Context)
}

type UsecaseCommand interface {
	PostRegister(ctx *gin.Context)
	PostLogin(ctx *gin.Context)
}

type RepositoryQuery interface {
	FindOneByID(ctx *gin.Context, id string) utils.Result
	FindOneByName(ctx *gin.Context, name string) utils.Result
}

type RepositoryCommand interface {
	Create(ctx *gin.Context, u models.User) utils.Result
	Save(ctx *gin.Context, u models.User) utils.Result
	FindPassword(ctx *gin.Context, u string) utils.FindPasswordResult
}
