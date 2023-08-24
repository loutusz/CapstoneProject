package usecases

import (
	"errors"
	"fmt"
	"login-api-jwt/bin/modules/messageprovider"
	"login-api-jwt/bin/pkg/databases"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// QueryUsecase implements messageprovider.UsecaseQuery interface
type QueryUsecase struct {
	MessageProviderRepositoryQuery messageprovider.RepositoryQuery
	ORM                            *databases.ORM
}

// NewQueryUsecase creates a new instance of QueryUsecase
func NewQueryUsecase(q messageprovider.RepositoryQuery, orm *databases.ORM) messageprovider.UsecaseQuery {
	return &QueryUsecase{
		MessageProviderRepositoryQuery: q,
		ORM:                            orm,
	}
}

// GetByID retrieves messageprovider data by ID and responds with the result
func (q QueryUsecase) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")

	// Call FindOneByID method to retrieve messageprovider data by ID
	ret := q.MessageProviderRepositoryQuery.FindOneByID(ctx, id)
	// If there was an error during query, abort with a Bad Request status
	if ret.DB.Error != nil {
		if errors.Is(ret.DB.Error, gorm.ErrRecordNotFound) {
			// If data is not found in the database, abort with status Unauthorized
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("MessageProvider with id %s not found", id)})
			return
		}

		ctx.AbortWithError(http.StatusBadRequest, ret.DB.Error)
		return
	}

	// Respond with retrieved messageprovider data in JSON format
	res := ret.Data
	ctx.JSON(http.StatusOK, res)
}

// GetAccess retrieves All messageprovider data and responds with the result
func (q QueryUsecase) GetAll(ctx *gin.Context) {
	ret := q.MessageProviderRepositoryQuery.FindAll(ctx)
	if ret.DB.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, ret.DB.Error)
		return
	}
	res := ret.Data
	ctx.JSON(http.StatusOK, res)
}

// GetAccess responds with a success message indicating messageprovider access
func (q QueryUsecase) GetAccess(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "messageprovider access success"})
}
