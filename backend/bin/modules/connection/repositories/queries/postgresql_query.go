package queries

import (
	"login-api-jwt/bin/modules/connection"
	"login-api-jwt/bin/modules/connection/models"
	"login-api-jwt/bin/pkg/databases"
	"login-api-jwt/bin/pkg/utils"

	"github.com/gin-gonic/gin"
)

// QueryRepository implements connection.RepositoryQuery interface
type QueryRepository struct {
	ORM *databases.ORM
}

// NewQueryRepository creates a new instance of QueryRepository
func NewQueryRepository(orm *databases.ORM) connection.RepositoryQuery {
	return &QueryRepository{
		ORM: orm,
	}
}

// FindOneByID retrieves a connection record from database by ID

func (q QueryRepository) FindAll(ctx *gin.Context, skip, limit int) utils.Result {
	var connectionsModel []models.Connection

	// Use ORM to find a connection record by ID
	r := q.ORM.DB.Offset(skip).Limit(limit).Find(&connectionsModel)

	// Prepare the result, including retrieved connection data and database operation result
	output := utils.Result{
		Data: connectionsModel,
		DB:   r,
	}
	return output

}
func (q QueryRepository) FindOneByID(ctx *gin.Context, connection_id string) utils.Result {
	var connectionModel models.Connection

	// Use ORM to find a connection record by ID
	r := q.ORM.DB.First(&connectionModel, "connection_id = ?", connection_id)
	// Prepare the result, including retrieved connection data and database operation result
	output := utils.Result{
		Data: connectionModel,
		DB:   r,
	}
	return output

}

func (q QueryRepository) FindOneByProjectID(ctx *gin.Context, project_id string) utils.Result {
	var connectionModel models.Connection

	// Use ORM to find a connection record by ID
	r := q.ORM.DB.First(&connectionModel, "project_id = ?", project_id)
	// Prepare the result, including retrieved connection data and database operation result
	output := utils.Result{
		Data: connectionModel,
		DB:   r,
	}
	return output

}

func (q QueryRepository) FindOneByMessageProviderID(ctx *gin.Context, message_provider_id string) utils.Result {
	var connectionModel models.Connection

	// Use ORM to find a connection record by ID
	r := q.ORM.DB.First(&connectionModel, "message_provider_id = ?", message_provider_id)
	// Prepare the result, including retrieved connection data and database operation result
	output := utils.Result{
		Data: connectionModel,
		DB:   r,
	}
	return output

}

func (q QueryRepository) CountData(ctx *gin.Context) utils.Result {
	var connectionModel models.Connection
	var count int64
	r := q.ORM.DB.Find(&connectionModel).Count(&count)

	output := utils.Result{
		Data: count,
		DB:   r,
	}
	return output
}
