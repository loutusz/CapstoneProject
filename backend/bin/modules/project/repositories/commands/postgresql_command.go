package queries

import (
	connectionModels "login-api-jwt/bin/modules/connection/models"
	"login-api-jwt/bin/modules/project"
	"login-api-jwt/bin/modules/project/models"
	"login-api-jwt/bin/pkg/databases"
	"login-api-jwt/bin/pkg/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CommandRepository implements project.RepositoryCommand interface
type CommandRepository struct {
	ORM *databases.ORM
}

// NewCommandRepository creates a new instance of CommandRepository
func NewCommandRepository(orm *databases.ORM) project.RepositoryCommand {
	return &CommandRepository{
		ORM: orm,
	}
}

// Create creates a new project record in database
func (c *CommandRepository) Create(ctx *gin.Context, p models.Project) utils.Result {
	// Use ORM to create a new project record in database
	r := c.ORM.DB.Create(&p)
	// Prepare the result, including project data and database operation result
	output := utils.Result{
		Data: p,
		DB:   r,
	}
	return output
}

// Save updates an existing project record in database
func (c *CommandRepository) Save(ctx *gin.Context, p models.Project) utils.Result {
	// Use ORM to update an existing project record in database
	r := c.ORM.DB.Save(&p)
	// Prepare the result, including project data and database operation result
	output := utils.Result{
		Data: p,
		DB:   r,
	}
	return output
}

func (c *CommandRepository) Updates(ctx *gin.Context, p models.Project) utils.Result {

	r := c.ORM.DB.Updates(&p)

	output := utils.Result{
		Data: p,
		DB:   r,
	}
	return output
}

func (c *CommandRepository) Delete(ctx *gin.Context, id string) utils.MultiDataResult {
	var projectModel models.Project
	var connectionModel connectionModels.Connection

	c.ORM.DB.First(&connectionModel, "project_id = ?", id)
	c.ORM.DB.First(&projectModel, "id = ?", id)

	connectionRecordset := c.ORM.DB.Delete(&connectionModel, "project_id = ?", id)
	projectRecordset := c.ORM.DB.Delete(&projectModel, "id = ?", id)

	result := struct {
		Project    models.Project
		Connection connectionModels.Connection
	}{
		Project:    projectModel,
		Connection: connectionModel,
	}

	output := utils.MultiDataResult{
		Data: result,
		DB:   []*gorm.DB{connectionRecordset, projectRecordset},
	}
	return output
}
