package usecases

import (
	"errors"
	"fmt"
	"login-api-jwt/bin/modules/user"
	"login-api-jwt/bin/pkg/databases"
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
	id := ctx.Param("id")

	// Call FindOneByID method to retrieve user data by ID
	ret := q.UserRepositoryQuery.FindOneByID(ctx, id)
	// If there was an error during query, abort with a Bad Request status
	if ret.DB.Error != nil {
		if errors.Is(ret.DB.Error, gorm.ErrRecordNotFound) {
			// If data is not found in the database, abort with status Unauthorized
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User with id %s not found", id)})
			return
		}

		ctx.AbortWithError(http.StatusBadRequest, ret.DB.Error)
		return
	}

	// Respond with retrieved user data in JSON format
	res := ret.Data
	ctx.JSON(http.StatusOK, res)
}

// GetAccess responds with a success message indicating user access
func (q QueryUsecase) GetAccess(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "user access success"})
}

// GetByName retrieves user data by name and responds with the result
func (q QueryUsecase) GetByName(ctx *gin.Context) {
	name := ctx.Param("name")
	fmt.Printf("name access %s", name)

	// Call FindOneByName method to retrieve user data by name
	ret := q.UserRepositoryQuery.FindOneByName(ctx, name)

	// Respond with retrieved user data in JSON format
	ctx.JSON(http.StatusOK, ret.Data)
}
