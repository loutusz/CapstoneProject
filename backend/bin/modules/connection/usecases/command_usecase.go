package usecases

import (
	"login-api-jwt/bin/modules/connection"
	"login-api-jwt/bin/modules/connection/models"
	"login-api-jwt/bin/pkg/databases"
	"login-api-jwt/bin/pkg/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CommandUsecase implements connection.UsecaseCommand interface
type CommandUsecase struct {
	ConnectionRepositoryCommand connection.RepositoryCommand
	ORM                         *databases.ORM
}

// NewCommandUsecase creates a new instance of CommandUsecase
func NewCommandUsecase(q connection.RepositoryCommand, orm *databases.ORM) connection.UsecaseCommand {
	return &CommandUsecase{
		ConnectionRepositoryCommand: q,
		ORM:                         orm,
	}
}

// PostRegister handles connection registration
func (q CommandUsecase) PostConnection(ctx *gin.Context) {
	var result utils.ResultResponse = utils.ResultResponse{
		Code:    http.StatusBadRequest,
		Data:    nil,
		Message: "Failed Post Connection",
		Status:  false,
	}
	var connectionModel models.Connection
	err := ctx.ShouldBind(&connectionModel)
	if err != nil {
		ctx.AbortWithStatusJSON(result.Code, result)
		return
	}

	// Generate a unique ID for connection
	connectionModel.ID = uuid.NewString()

	// Create connection record in the database
	r := q.ConnectionRepositoryCommand.Create(ctx, connectionModel)
	if r.DB.Error != nil {

		if strings.Contains(r.DB.Error.Error(), "insert or update on table \"connections\" violates foreign key constraint \"connections_project_id_fkey\"") {
			// If data user id not valid return message "user id not valid"
			result.Message = "project id not valid"
			ctx.AbortWithStatusJSON(result.Code, result)
			return
		}

		if strings.Contains(r.DB.Error.Error(), "insert or update on table \"connections\" violates foreign key constraint \"connections_message_provider_id_fkey\"") {
			// If data user id not valid return message "user id not valid"
			result.Message = "message_provider id not valid"
			ctx.AbortWithStatusJSON(result.Code, result)
			return
		}
		result.Code = http.StatusInternalServerError
		ctx.AbortWithError(http.StatusInternalServerError, r.DB.Error)
		return
	}

	// Response data for successful registration
	connectionRegisterResponse := models.PostConnectionResponse{
		ID:                  connectionModel.ID,
		Message_provider_id: connectionModel.Message_provider_id,
		Project_id:          connectionModel.Project_id,
	}

	// Save connection record again after successful registration
	r = q.ConnectionRepositoryCommand.Save(ctx, connectionModel)

	// Check if an error occurred while saving
	if r.DB.Error != nil {
		// If there was an error, return Internal Server Error with error message
		ctx.AbortWithError(http.StatusInternalServerError, r.DB.Error)
		return
	}
	// If connection record was successfully saved, respond with connection's registration data
	ctx.JSON(http.StatusOK, connectionRegisterResponse)
}

func (q CommandUsecase) PutConnection(ctx *gin.Context) {
	connectionID := ctx.Param("id")
	var connectionModel models.Connection
	err := ctx.ShouldBind(&connectionModel)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	connectionModel.ID = connectionID

	// Response data for successful registration
	Response := connectionModel

	r := q.ConnectionRepositoryCommand.Updates(ctx, Response)
	if r.DB.Error != nil {
		// If there was an error, return Internal Server Error with error message
		ctx.AbortWithError(http.StatusInternalServerError, r.DB.Error)
		return
	}

	if r.DB.RowsAffected == 0 {
		// If there was an error, return Internal Server Error with error message
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Connection ID not available"})
		return
	}
	// If messageprovider record was successfully saved, respond with messageprovider's registration data
	ctx.JSON(http.StatusOK, Response)

}
