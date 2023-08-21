package usecases

import (
	"errors"
	"fmt"
	"login-api-jwt/bin/modules/project"
	"login-api-jwt/bin/pkg/databases"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// QueryUsecase implements project.UsecaseQuery interface
type QueryUsecase struct {
	ProjectRepositoryQuery project.RepositoryQuery
	ORM                    *databases.ORM
}

// NewQueryUsecase creates a new instance of QueryUsecase
func NewQueryUsecase(q project.RepositoryQuery, orm *databases.ORM) project.UsecaseQuery {
	return &QueryUsecase{
		ProjectRepositoryQuery: q,
		ORM:                    orm,
	}
}

// GetByID retrieves project data by ID and responds with the result
func (q QueryUsecase) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")

	// Call FindOneByID method to retrieve project data by ID
	ret := q.ProjectRepositoryQuery.FindOneByID(ctx, id)
	// If there was an error during query, abort with a Bad Request status
	if ret.DB.Error != nil {
		if errors.Is(ret.DB.Error, gorm.ErrRecordNotFound) {
			// If data is not found in the database, abort with status Unauthorized
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Project with id %s not found", id)})
			return
		}

		ctx.AbortWithError(http.StatusBadRequest, ret.DB.Error)
		return
	}

	// Respond with retrieved project data in JSON format
	res := ret.Data
	ctx.JSON(http.StatusOK, res)
}

// GetAccess retrieves All project data and responds with the result
func (q QueryUsecase) GetAll(ctx *gin.Context) {
	ret := q.ProjectRepositoryQuery.FindAll(ctx)
	if ret.DB.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, ret.DB.Error)
		return
	}
	res := ret.Data
	ctx.JSON(http.StatusOK, res)
}

// GetAccess responds with a success message indicating project access
func (q QueryUsecase) GetAccess(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "project access success"})
}

// // GetByName retrieves project data by name and responds with the result
// func (q QueryUsecase) GetByName(ctx *gin.Context) {
// 	name := ctx.Param("name")
// 	fmt.Printf("name access %s", name)

// 	// Call FindOneByName method to retrieve project data by name
// 	ret := q.ProjectRepositoryQuery.FindOneByName(ctx, name)

// 	// Respond with retrieved project data in JSON format
// 	ctx.JSON(http.StatusOK, ret.Data)
// }
