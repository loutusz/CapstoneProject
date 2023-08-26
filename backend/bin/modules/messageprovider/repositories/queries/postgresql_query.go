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

func (q QueryRepository) FindAll(ctx *gin.Context, skip, limit int) utils.Result {
	var messageprovidersModel []models.MessageProvider

	// Use ORM to find a messageprovider record by ID
	r := q.ORM.DB.Offset(skip).Limit(limit).Find(&messageprovidersModel)

	// Prepare the result, including retrieved messageprovider data and database operation result
	output := utils.Result{
		Data: messageprovidersModel,
		DB:   r,
	}
	return output

}

func (q QueryRepository) FindOneByID(ctx *gin.Context, message_id string) utils.Result {
	var messageproviderModel models.MessageProvider

	// Use ORM to find a messageprovider record by ID
	r := q.ORM.DB.First(&messageproviderModel, "message_provider_id = ?", message_id)
	// Prepare the result, including retrieved messageprovider data and database operation result
	output := utils.Result{
		Data: messageproviderModel,
		DB:   r,
	}
	return output

}
func (q QueryRepository) CountData(ctx *gin.Context) utils.Result {
	var messageproviderModel models.MessageProvider
	var count int64
	r := q.ORM.DB.Find(&messageproviderModel).Count(&count)

	output := utils.Result{
		Data: count,
		DB:   r,
	}
	return output
}

func (q QueryRepository) FindByUserID(ctx *gin.Context, user_id string, skip, limit int) utils.Result {
	var messageprovidersModel []models.MessageProvider

	// Use ORM to find a messageprovider record by ID
	r := q.ORM.DB.Where("message_provider_user_id = ?", user_id).Offset(skip).Limit(limit).Find(&messageprovidersModel)

	// Prepare the result, including retrieved messageprovider data and database operation result
	output := utils.Result{
		Data: messageprovidersModel,
		DB:   r,
	}
	return output

}
