package usecases

import (
	"login-api-jwt/bin/modules/messageprovider"
	"login-api-jwt/bin/modules/messageprovider/models"
	"login-api-jwt/bin/pkg/databases"
	"login-api-jwt/bin/pkg/utils"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
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
	var messageProviderModel models.MessageProvider
	err := ctx.ShouldBind(&messageProviderModel)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		return
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	token, err := utils.ValidateUserJWTToToken(tokenString)

	if err != nil {
		if err.Error() == "invalid token" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid claims"})
	}
	messageProviderModel.MessageProviderUserID = claims["id"].(string)

	// Generate a unique MessageProviderID for messageprovider
	messageProviderModel.MessageProviderID = uuid.NewString()

	// Create messageprovider record in the database

	r := q.MessageProviderRepositoryCommand.Create(ctx, messageProviderModel)
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
	messageproviderRegisterResponse := messageProviderModel

	// Save messageprovider record again after successful registration
	r = q.MessageProviderRepositoryCommand.Save(ctx, messageProviderModel)

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

	messageProviderModel.MessageProviderID = messageProviderID

	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		return
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	token, err := utils.ValidateUserJWTToToken(tokenString)

	if err != nil {
		if err.Error() == "invalid token" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid claims"})
	}
	messageProviderModel.MessageProviderUserID = claims["id"].(string)

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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Message Provider MessageProviderID not available"})
		return
	}
	// If messageprovider record was successfully saved, respond with messageprovider's registration data
	ctx.JSON(http.StatusOK, Response)

}

func (q CommandUsecase) DeleteMessageProvider(ctx *gin.Context) {
	var result utils.ResultResponse = utils.ResultResponse{
		Code:    http.StatusBadRequest,
		Data:    nil,
		Message: "Failed Delete MessageProvider",
		Status:  false,
	}

	var id string = ctx.Param("id")

	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		return
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	_, err := utils.ValidateUserJWTToToken(tokenString)

	if err != nil {
		if err.Error() == "invalid token" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	deletedMessageProvider := q.MessageProviderRepositoryCommand.Delete(ctx, id)
	if deletedMessageProvider.DB[0].Error != nil {
		result.Code = http.StatusInternalServerError
		ctx.AbortWithStatusJSON(result.Code, result)
		return
	}

	if deletedMessageProvider.DB[1].Error != nil {
		result.Code = http.StatusInternalServerError
		ctx.AbortWithStatusJSON(result.Code, result)
		return
	}

	if deletedMessageProvider.DB[1].RowsAffected == 0 {
		// If there was an error, return Internal Server Error with error message
		result.Code = http.StatusBadRequest
		result.Message = "messageProvider not found"
		ctx.AbortWithStatusJSON(result.Code, result)
		return
	}
	result = utils.ResultResponse{
		Code:    http.StatusOK,
		Data:    deletedMessageProvider.Data,
		Message: "Success Delete MessageProvider",
		Status:  true,
	}
	ctx.JSON(result.Code, result)
}
