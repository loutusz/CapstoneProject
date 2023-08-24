package project

import (
	"login-api-jwt/bin/modules/project/models"
	"login-api-jwt/bin/pkg/utils"

	"github.com/gin-gonic/gin"
)

type UsecaseQuery interface {
	GetByID(ctx *gin.Context)
	GetAccess(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	// GetByName(ctx *gin.Context)
}

type UsecaseCommand interface {
	PostProject(ctx *gin.Context)
	PutProject(ctx *gin.Context)
}

type RepositoryQuery interface {
	FindAll(ctx *gin.Context, skip, limit int) utils.Result
	FindOneByID(ctx *gin.Context, id string) utils.Result
	CountData(ctx *gin.Context) utils.Result
}

type RepositoryCommand interface {
	Create(ctx *gin.Context, u models.Project) utils.Result
	Save(ctx *gin.Context, u models.Project) utils.Result
	Updates(ctx *gin.Context, u models.Project) utils.Result
}
