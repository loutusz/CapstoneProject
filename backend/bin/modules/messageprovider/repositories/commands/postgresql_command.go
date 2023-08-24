package queries

import (
	"login-api-jwt/bin/modules/messageprovider"
	"login-api-jwt/bin/modules/messageprovider/models"
	"login-api-jwt/bin/pkg/databases"
	"login-api-jwt/bin/pkg/utils"

	"github.com/gin-gonic/gin"
)

// CommandRepository implements messageprovider.RepositoryCommand interface
type CommandRepository struct {
	ORM *databases.ORM
}

// NewCommandRepository creates a new instance of CommandRepository
func NewCommandRepository(orm *databases.ORM) messageprovider.RepositoryCommand {
	return &CommandRepository{
		ORM: orm,
	}
}

// Create creates a new messageprovider record in database
func (c *CommandRepository) Create(ctx *gin.Context, p models.MessageProvider) utils.Result {
	// Use ORM to create a new messageprovider record in database
	r := c.ORM.DB.Create(&p)
	// Prepare the result, including messageprovider data and database operation result
	output := utils.Result{
		Data: p,
		DB:   r,
	}
	return output
}

// Save updates an existing messageprovider record in database
func (c *CommandRepository) Save(ctx *gin.Context, p models.MessageProvider) utils.Result {
	// Use ORM to update an existing messageprovider record in database
	r := c.ORM.DB.Save(&p)
	// Prepare the result, including messageprovider data and database operation result
	output := utils.Result{
		Data: p,
		DB:   r,
	}
	return output
}

func (c *CommandRepository) Updates(ctx *gin.Context, m models.MessageProvider) utils.Result {

	r := c.ORM.DB.Updates(&m)

	output := utils.Result{
		Data: m,
		DB:   r,
	}
	return output
}
