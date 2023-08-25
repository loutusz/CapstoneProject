package queries

import (
	connectionModels "login-api-jwt/bin/modules/connection/models"
	"login-api-jwt/bin/modules/messageprovider"
	"login-api-jwt/bin/modules/messageprovider/models"
	"login-api-jwt/bin/pkg/databases"
	"login-api-jwt/bin/pkg/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func (c *CommandRepository) Delete(ctx *gin.Context, id string) utils.MultiDataResult {
	var messageProviderModel models.MessageProvider
	var connectionModel connectionModels.Connection

	c.ORM.DB.First(&connectionModel, "message_provider_id = ?", id)
	c.ORM.DB.First(&messageProviderModel, "id = ?", id)

	connectionRecordset := c.ORM.DB.Delete(&connectionModel, "message_provider_id = ?", id)
	messageProviderRecordset := c.ORM.DB.Delete(&messageProviderModel, "id = ?", id)

	result := struct {
		MessageProvider models.MessageProvider
		Connection      connectionModels.Connection
	}{
		MessageProvider: messageProviderModel,
		Connection:      connectionModel,
	}

	output := utils.MultiDataResult{
		Data: result,
		DB:   []*gorm.DB{connectionRecordset, messageProviderRecordset},
	}
	return output
}
