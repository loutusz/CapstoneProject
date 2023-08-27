package queries

import (
	"login-api-jwt/bin/modules/user"
	"login-api-jwt/bin/modules/user/models"
	"login-api-jwt/bin/pkg/databases"
	"login-api-jwt/bin/pkg/utils"

	"github.com/gin-gonic/gin"
)

// QueryRepository implements user.RepositoryQuery interface
type QueryRepository struct {
	ORM *databases.ORM
}

// NewQueryRepository creates a new instance of QueryRepository
func NewQueryRepository(orm *databases.ORM) user.RepositoryQuery {
	return &QueryRepository{
		ORM: orm,
	}
}

// FindOneByID retrieves a user record from database by ID
func (q QueryRepository) FindOneByID(ctx *gin.Context, id string) utils.Result {
	var userModel models.User

	// Use ORM to find a user record by ID
	r := q.ORM.DB.First(&userModel, "id = ?", id)
	// Prepare the result, including retrieved user data and database operation result
	output := utils.Result{
		Data: userModel,
		DB:   r,
	}
	return output

}

// FindOneByName retrieves a user record from database by name
func (q QueryRepository) FindOneByName(ctx *gin.Context, name string) utils.Result {
	var userModel models.User

	// Use ORM to find a user record by name
	r := q.ORM.DB.First(&userModel, "name = ?", name)
	// Prepare the result, including retrieved user data and database operation result
	output := utils.Result{
		Data: userModel,
		DB:   r,
	}
	return output

}
