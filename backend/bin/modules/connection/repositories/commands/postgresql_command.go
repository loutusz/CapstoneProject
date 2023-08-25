package queries

import (
	"login-api-jwt/bin/modules/connection"
	"login-api-jwt/bin/modules/connection/models"
	"login-api-jwt/bin/pkg/databases"
	"login-api-jwt/bin/pkg/utils"

	"github.com/gin-gonic/gin"
)

// CommandRepository implements connection.RepositoryCommand interface
type CommandRepository struct {
	ORM *databases.ORM
}

// NewCommandRepository creates a new instance of CommandRepository
func NewCommandRepository(orm *databases.ORM) connection.RepositoryCommand {
	return &CommandRepository{
		ORM: orm,
	}
}

// Create creates a new connection record in database
func (c *CommandRepository) Create(ctx *gin.Context, p models.Connection) utils.Result {
	// Use ORM to create a new connection record in database
	r := c.ORM.DB.Create(&p)
	// Prepare the result, including connection data and database operation result
	output := utils.Result{
		Data: p,
		DB:   r,
	}
	return output
}

// Save updates an existing connection record in database
func (c *CommandRepository) Save(ctx *gin.Context, p models.Connection) utils.Result {
	// Use ORM to update an existing connection record in database
	r := c.ORM.DB.Save(&p)
	// Prepare the result, including connection data and database operation result
	output := utils.Result{
		Data: p,
		DB:   r,
	}
	return output
}

func (c *CommandRepository) Updates(ctx *gin.Context, p models.Connection) utils.Result {

	r := c.ORM.DB.Updates(&p)

	output := utils.Result{
		Data: p,
		DB:   r,
	}
	return output
}

func (c *CommandRepository) Delete(ctx *gin.Context, id string) utils.Result {
	var connectionModel models.Connection
	recordset := c.ORM.DB.Delete(&connectionModel, "id = ?", id)

	output := utils.Result{
		Data: connectionModel,
		DB:   recordset,
	}
	return output
}
