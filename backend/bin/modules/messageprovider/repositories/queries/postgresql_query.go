package queries

import (
	"login-api-jwt/bin/modules/messageprovider"
	"login-api-jwt/bin/modules/messageprovider/models"
	"login-api-jwt/bin/pkg/databases"
	"login-api-jwt/bin/pkg/utils"

	"github.com/gin-gonic/gin"
)

// QueryRepository implements messageprovider.RepositoryQuery interface
type QueryRepository struct {
	ORM *databases.ORM
}

// NewQueryRepository creates a new instance of QueryRepository
func NewQueryRepository(orm *databases.ORM) messageprovider.RepositoryQuery {
	return &QueryRepository{
		ORM: orm,
	}
}

// FindOneByID retrieves a messageprovider record from database by ID

func (q QueryRepository) FindAll(ctx *gin.Context) utils.Result {
	var messageprovidersModel []models.MessageProvider

	// Use ORM to find a messageprovider record by ID
	r := q.ORM.DB.Find(&messageprovidersModel)

	// Prepare the result, including retrieved messageprovider data and database operation result
	output := utils.Result{
		Data: messageprovidersModel,
		DB:   r,
	}
	return output

}
func (q QueryRepository) FindOneByID(ctx *gin.Context, id string) utils.Result {
	var messageproviderModel models.MessageProvider

	// Use ORM to find a messageprovider record by ID
	r := q.ORM.DB.First(&messageproviderModel, "id = ?", id)
	// Prepare the result, including retrieved messageprovider data and database operation result
	output := utils.Result{
		Data: messageproviderModel,
		DB:   r,
	}
	return output

}

// // FindOneByName retrieves a messageprovider record from database by name
// func (q QueryRepository) FindOneByName(ctx *gin.Context, name string) utils.Result {
// 	var messageproviderModel models.MessageProvider

// 	// Use ORM to find a messageprovider record by name
// 	r := q.ORM.DB.First(&messageproviderModel, "name = ?", name)
// 	// Prepare the result, including retrieved messageprovider data and database operation result
// 	output := utils.Result{
// 		Data: messageproviderModel,
// 		DB:   r,
// 	}
// 	return output

// }
