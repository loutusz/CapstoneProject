package queries

import (
	"login-api-jwt/bin/modules/project"
	"login-api-jwt/bin/modules/project/models"
	"login-api-jwt/bin/pkg/databases"
	"login-api-jwt/bin/pkg/utils"

	"github.com/gin-gonic/gin"
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
