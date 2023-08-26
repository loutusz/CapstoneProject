package queries

import (
	connectionModels "login-api-jwt/bin/modules/connection/models"
	messageProviderModels "login-api-jwt/bin/modules/messageprovider/models"
	"login-api-jwt/bin/modules/project"
	"login-api-jwt/bin/modules/project/models"

	"login-api-jwt/bin/pkg/databases"
	"login-api-jwt/bin/pkg/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func (q QueryRepository) FindOneByID(ctx *gin.Context, project_id string) utils.Result {
	var projectModel models.Project

	// Use ORM to find a project record by ID
	r := q.ORM.DB.First(&projectModel, "project_id = ?", project_id)
	// Prepare the result, including retrieved project data and database operation result
	output := utils.Result{
		Data: projectModel,
		DB:   r,
	}
	return output

}

func (q QueryRepository) FindConnectedOneByID(ctx *gin.Context, project_id string) utils.Result {
	var projectModel models.Project
	// Define the structs for the connection and provider
	var connectionModel connectionModels.Connection
	var providerModel messageProviderModels.MessageProvider

	// Use ORM to find a project record by ID
	r := q.ORM.DB.First(&projectModel, "project_id = ?", project_id)
	connector := struct {
		Data    interface{}
		Message string
	}{
		Message: "found",
		Data:    &connectionModel,
	}

	provider := struct {
		Data    interface{}
		Message string
	}{
		Message: "found",
		Data:    &providerModel,
	}

	if q.ORM.DB.First(&connectionModel, "connection_project_id = ?", project_id).Error == gorm.ErrRecordNotFound {
		connector.Message = "record not found"

	}

	// Use ORM to find a provider record by provider ID
	if q.ORM.DB.First(&providerModel, "message_provider_id = ?", connectionModel.ConnectionMessageProviderID).Error == gorm.ErrRecordNotFound {
		// Handle the case when provider data is not found
		provider.Message = "record not found"
	}

	// Prepare the result
	output := utils.Result{
		Data: struct {
			Project   models.Project
			Connector interface{}
			Provider  interface{}
		}{
			Project:   projectModel,
			Connector: connector,
			Provider:  provider,
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

func (q QueryRepository) FindByUserID(ctx *gin.Context, user_id string, skip, limit int) utils.Result {
	var projectsModel []models.Project

	// Use ORM to find a project record by ID
	r := q.ORM.DB.Where("project_user_id = ?", user_id).Offset(skip).Limit(limit).Find(&projectsModel)

	// Prepare the result, including retrieved project data and database operation result
	output := utils.Result{
		Data: projectsModel,
		DB:   r,
	}
	return output

}

func (q QueryRepository) FindConnectedByUserID(ctx *gin.Context, user_id string, skip, limit int) utils.Result {
	// var projectInfo []struct {
	// 	Project         models.Project
	// 	Connection      connectionModels.Connection
	// 	MessageProvider messageProviderModels.MessageProvider
	// }

	var projectsInfo []map[string]interface{}

	// Use ORM to find project records by user ID with LEFT JOIN on connections and message_providers
	r := q.ORM.DB.
		Table("projects").
		Select("projects.*, connections.*, message_providers.*").
		Joins("LEFT JOIN connections ON connections.connection_project_id = projects.project_id").
		Joins("LEFT JOIN message_providers ON message_providers.message_provider_id = connections.connection_message_provider_id").
		Where("projects.project_user_id = ?", user_id).
		Offset(skip).
		Limit(limit).
		Scan(&projectsInfo)

	// Prepare the result, including retrieved project data and database operation result
	output := utils.Result{
		Data: projectsInfo,
		DB:   r,
	}
	return output
}
