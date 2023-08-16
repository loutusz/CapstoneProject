package queries

import (
	"login-api-jwt/bin/modules/user"
	"login-api-jwt/bin/modules/user/models"
	"login-api-jwt/bin/pkg/databases"
	"login-api-jwt/bin/pkg/utils"

	"github.com/gin-gonic/gin"
)

// CommandRepository implements user.RepositoryCommand interface
type CommandRepository struct {
	ORM *databases.ORM
}

// NewCommandRepository creates a new instance of CommandRepository
func NewCommandRepository(orm *databases.ORM) user.RepositoryCommand {
	return &CommandRepository{
		ORM: orm,
	}
}

// Create creates a new user record in database
func (c *CommandRepository) Create(ctx *gin.Context, u models.User) utils.Result {
	// Use ORM to create a new user record in database
	r := c.ORM.DB.Create(&u)
	// Prepare the result, including user data and database operation result
	output := utils.Result{
		Data: u,
		DB:   r,
	}
	return output
}

// Save updates an existing user record in database
func (c *CommandRepository) Save(ctx *gin.Context, u models.User) utils.Result {
	// Use ORM to update an existing user record in database
	r := c.ORM.DB.Save(&u)
	// Prepare the result, including user data and database operation result
	output := utils.Result{
		Data: u,
		DB:   r,
	}
	return output
}

// FindPassword retrieves password for a user by username
func (c *CommandRepository) FindPassword(ctx *gin.Context, u string) utils.FindPasswordResult {
	var userModel models.User
	userModel.Username = u
	// Use ORM to find user record by username
	r := c.ORM.DB.First(&userModel, "username = ?", u)
	// Prepare the result, including retrieved password, user data, and  database operation result
	output := utils.FindPasswordResult{
		Password: userModel.Password,
		Data:     userModel,
		DB:       r,
	}
	return output
}
