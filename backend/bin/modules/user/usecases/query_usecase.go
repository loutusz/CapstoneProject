package usecases

import (
	"errors"
	"fmt"
	"login-api-jwt/bin/modules/user"
	"login-api-jwt/bin/pkg/databases"
	"login-api-jwt/bin/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// QueryUsecase implements user.UsecaseQuery interface
type QueryUsecase struct {
	UserRepositoryQuery user.RepositoryQuery
	ORM                 *databases.ORM
}

// NewQueryUsecase creates a new instance of QueryUsecase
func NewQueryUsecase(q user.RepositoryQuery, orm *databases.ORM) user.UsecaseQuery {
	return &QueryUsecase{
		UserRepositoryQuery: q,
		ORM:                 orm,
	}
}

// GetByID retrieves user data by ID and responds with the result
func (q QueryUsecase) GetByID(ctx *gin.Context) {
	var result utils.ResultResponse = utils.ResultResponse{
		Code:    http.StatusBadRequest,
		Data:    nil,
		Message: "Failed Get Data User",
		Status:  false,
	}
	id := ctx.Param("id")

	// Call FindOneByID method to retrieve user data by ID
	userData := q.UserRepositoryQuery.FindOneByID(ctx, id)
	// If there was an error during query, abort with a Bad Request status
	if userData.DB.Error != nil {
		if errors.Is(userData.DB.Error, gorm.ErrRecordNotFound) {
			if errors.Is(userData.DB.Error, gorm.ErrRecordNotFound) {
				result.Code = http.StatusNotFound
				result.Message = "Data Not Found"
				ctx.AbortWithStatusJSON(result.Code, result)
				return
			}
			// If data is not found in the database, abort with status Unauthorized
			ctx.AbortWithStatusJSON(http.StatusNotFound, result)
			return
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, result)
		return
	}

	result = utils.ResultResponse{
		Code:    http.StatusOK,
		Data:    userData.Data,
		Message: "Success Get Data User",
		Status:  true,
	}

	// Respond with retrieved user data in JSON format
	ctx.JSON(http.StatusOK, result)
}

// GetAccess responds with a success message indicating user access
func (q QueryUsecase) GetAccess(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "user access success"})
}

// GetByName retrieves user data by name and responds with the result
func (q QueryUsecase) GetByUsername(ctx *gin.Context) {
	var result utils.ResultResponse = utils.ResultResponse{
		Code:    http.StatusBadRequest,
		Data:    nil,
		Message: "Failed Get Data User",
		Status:  false,
	}
	username := ctx.Param("username")
	fmt.Printf("username access %s", username)

	// Call FindOneByUsername method to retrieve user data by username
	userData := q.UserRepositoryQuery.FindOneByUsername(ctx, username)
	if userData.DB.Error != nil {
		if errors.Is(userData.DB.Error, gorm.ErrRecordNotFound) {
			result.Code = http.StatusNotFound
			result.Message = "Data Not Found"
			ctx.AbortWithStatusJSON(result.Code, result)
			return
		}

		ctx.AbortWithStatusJSON(result.Code, result)
		return
	}

	result = utils.ResultResponse{
		Code:    http.StatusOK,
		Data:    userData.Data,
		Message: "Success Get Data User",
		Status:  true,
	}
	// Respond with retrieved user data in JSON format
	ctx.JSON(http.StatusOK, result)
}
