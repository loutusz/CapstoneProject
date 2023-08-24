package usecases

import (
	"login-api-jwt/bin/modules/messageprovider"
	"login-api-jwt/bin/modules/messageprovider/models"
	"login-api-jwt/bin/pkg/databases"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CommandUsecase implements messageprovider.UsecaseCommand interface
type CommandUsecase struct {
	MessageProviderRepositoryCommand messageprovider.RepositoryCommand
	ORM                              *databases.ORM
}

// NewCommandUsecase creates a new instance of CommandUsecase
func NewCommandUsecase(q messageprovider.RepositoryCommand, orm *databases.ORM) messageprovider.UsecaseCommand {
	return &CommandUsecase{
		MessageProviderRepositoryCommand: q,
		ORM:                              orm,
	}
}

// PostRegister handles messageprovider registration
func (q CommandUsecase) PostMessageProvider(ctx *gin.Context) {
	var messageproviderModel models.MessageProvider
	err := ctx.ShouldBind(&messageproviderModel)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	// Generate a unique ID for messageprovider
	messageproviderModel.ID = uuid.NewString()

	// Create messageprovider record in the database

	r := q.MessageProviderRepositoryCommand.Create(ctx, messageproviderModel)
	if r.DB.Error != nil {
		if strings.Contains(r.DB.Error.Error(), "insert or update on table \"message_providers\" violates foreign key constraint \"message_providers_messageProvider_id_fkey\"") {
			// If data is already found, abort with status "email or messageProvidername already used"
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user id not valid"})
			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, r.DB.Error)
		return
	}

	// Response data for successful registration
	messageproviderRegisterResponse := messageproviderModel

	// Save messageprovider record again after successful registration
	r = q.MessageProviderRepositoryCommand.Save(ctx, messageproviderModel)

	// Check if an error occurred while saving
	if r.DB.Error != nil {
		// If there was an error, return Internal Server Error with error message
		ctx.AbortWithError(http.StatusInternalServerError, r.DB.Error)
		return
	}

	// If messageprovider record was successfully saved, respond with messageprovider's registration data
	ctx.JSON(http.StatusOK, messageproviderRegisterResponse)
}

func (q CommandUsecase) PutMessageProvider(ctx *gin.Context) {
	messageProviderID := ctx.Param("id")
	var messageProviderModel models.MessageProvider
	err := ctx.ShouldBind(&messageProviderModel)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "input field not valid")
	}

	messageProviderModel.ID = messageProviderID

	// Response data for successful registration
	Response := messageProviderModel

	r := q.MessageProviderRepositoryCommand.Updates(ctx, Response)
	if r.DB.Error != nil {
		// If there was an error, return Internal Server Error with error message
		ctx.AbortWithError(http.StatusInternalServerError, r.DB.Error)
		return
	}

	if r.DB.RowsAffected == 0 {
		// If there was an error, return Internal Server Error with error message
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Message Provider ID not available"})
		return
	}
	// If messageprovider record was successfully saved, respond with messageprovider's registration data
	ctx.JSON(http.StatusOK, Response)

}
