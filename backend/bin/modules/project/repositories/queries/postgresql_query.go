package queries

import (
	connectionModels "login-api-jwt/bin/modules/connection/models"
	messageProviderModels "login-api-jwt/bin/modules/messageprovider/models"
	"login-api-jwt/bin/modules/project"
	"login-api-jwt/bin/modules/project/models"

	"login-api-jwt/bin/pkg/databases"
	"login-api-jwt/bin/pkg/utils"

	"github.com/gin-gonic/gin"
)

// QueryRepository implements project.RepositoryQuery interface
type QueryRepository struct {
	ORM *databases.ORM
}

// NewQueryRepository creates a new instance of QueryRepository
func NewQueryRepository(orm *databases.ORM) project.RepositoryQuery {
	return &QueryRepository{
		ORM: orm,
	}
}

// FindOneByID retrieves a project record from database by ID

func (q QueryRepository) FindAll(ctx *gin.Context, skip, limit int) utils.Result {
	var projectsModel []models.Project

	// Use ORM to find a project record by ID
	r := q.ORM.DB.Offset(skip).Limit(limit).Find(&projectsModel)

	// Prepare the result, including retrieved project data and database operation result
	output := utils.Result{
		Data: projectsModel,
		DB:   r,
	}
	return output

}

func (q QueryRepository) FindOneByID(ctx *gin.Context, id string) utils.Result {
	var projectModel models.Project

	// Use ORM to find a project record by ID
	r := q.ORM.DB.First(&projectModel, "id = ?", id)
	// Prepare the result, including retrieved project data and database operation result
	output := utils.Result{
		Data: projectModel,
		DB:   r,
	}
	return output

}

func (q QueryRepository) FindConnectedOneByID(ctx *gin.Context, id string) utils.Result {
	var projectModel models.Project
	// Define the structs for the connection and provider
	var connectionModel connectionModels.Connection
	var providerModel messageProviderModels.MessageProvider

	// Use ORM to find a project record by ID
	r := q.ORM.DB.First(&projectModel, "id = ?", id)

	// Use ORM to find a project record by ID
	q.ORM.DB.First(&connectionModel, "project_id = ?", id)
	q.ORM.DB.First(&providerModel, "id = ?", connectionModel.Message_provider_id)

	// Prepare the result
	output := utils.Result{
		Data: struct {
			Project   models.Project
			Connector connectionModels.Connection
			Provider  messageProviderModels.MessageProvider
		}{
			Project:   projectModel,
			Connector: connectionModel,
			Provider:  providerModel,
		},
		DB: r,
	}
	return output
}

func (q QueryRepository) CountData(ctx *gin.Context) utils.Result {
	var projectModel models.Project
	var count int64
	r := q.ORM.DB.Find(&projectModel).Count(&count)

	output := utils.Result{
		Data: count,
		DB:   r,
	}
	return output
}

func (q QueryRepository) FindByUserID(ctx *gin.Context, id string, skip, limit int) utils.Result {
	var projectsModel []models.Project

	// Use ORM to find a project record by ID
	r := q.ORM.DB.Where("user_id = ?", id).Offset(skip).Limit(limit).Find(&projectsModel)

	// Prepare the result, including retrieved project data and database operation result
	output := utils.Result{
		Data: projectsModel,
		DB:   r,
	}
	return output

}

func (q QueryRepository) FindConnectedByUserID(ctx *gin.Context, id string, skip, limit int) utils.Result {
	var projectsModel []models.Project

	// Use ORM to find a project record by ID
	r := q.ORM.DB.Where("user_id = ?", id).Offset(skip).Limit(limit).Find(&projectsModel)

	// Prepare the result, including retrieved project data and database operation result
	output := utils.Result{
		Data: projectsModel,
		DB:   r,
	}
	return output

}
