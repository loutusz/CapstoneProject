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
	var result = utils.ResultResponse{
		Code:    http.StatusOK,
		Data:    nil,
		Message: "Failed Post Message Provider",
		Status:  false,
	}
	var messageProviderModel models.MessageProvider
	err := ctx.ShouldBind(&messageProviderModel)
	if err != nil {
		ctx.AbortWithStatusJSON(result.Code, result)
	}

	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		result.Code = http.StatusUnauthorized
		result.Message = "Token Required"
		ctx.AbortWithStatusJSON(result.Code, result)
		return
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	token, err := utils.ValidateUserJWTToToken(tokenString)

	if err != nil {
		if err.Error() == "invalid token" {
			result.Message = "Token is Expired"
			ctx.AbortWithStatusJSON(result.Code, result)
		}
		result.Code = http.StatusInternalServerError
		ctx.AbortWithStatusJSON(result.Code, result)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		result.Message = "Invalid claims"
		ctx.AbortWithStatusJSON(result.Code, result)
	}
	messageProviderModel.User_id = claims["id"].(string)

	// Generate a unique ID for messageprovider
	messageProviderModel.ID = uuid.NewString()

	// Create messageprovider record in the database

	r := q.MessageProviderRepositoryCommand.Create(ctx, messageProviderModel)
	if r.DB.Error != nil {
		if strings.Contains(r.DB.Error.Error(), "insert or update on table \"message_providers\" violates foreign key constraint \"message_providers_messageProvider_id_fkey\"") {
			// If data is already found, abort with status "email or messageProvidername already used"
			result.Message = "User id not valid"
			ctx.AbortWithStatusJSON(result.Code, result)
			return
		}
		result.Code = http.StatusInternalServerError
		ctx.AbortWithStatusJSON(result.Code, result)
		return
	}

	// Response data for successful registration
	messageproviderRegisterResponse := messageProviderModel

	// Save messageprovider record again after successful registration
	r = q.MessageProviderRepositoryCommand.Save(ctx, messageProviderModel)

	// Check if an error occurred while saving
	if r.DB.Error != nil {
		// If there was an error, return Internal Server Error with error message
		result.Code = http.StatusInternalServerError
		ctx.AbortWithStatusJSON(result.Code, result)
		return
	}

	result = utils.ResultResponse{
		Code:    http.StatusOK,
		Data:    messageproviderRegisterResponse,
		Message: "Success Register Data Message Provider",
		Status:  true,
	}
	// If messageprovider record was successfully saved, respond with messageprovider's registration data
	ctx.JSON(http.StatusOK, messageproviderRegisterResponse)
}

func (q CommandUsecase) PutMessageProvider(ctx *gin.Context) {
	var result = utils.ResultResponse{
		Code:    http.StatusOK,
		Data:    nil,
		Message: "Failed Post Message Provider",
		Status:  false,
	}

	messageProviderID := ctx.Param("id")
	var messageProviderModel models.MessageProvider
	err := ctx.ShouldBind(&messageProviderModel)
	if err != nil {
		ctx.AbortWithStatusJSON(result.Code, result)
	}

	messageProviderModel.ID = messageProviderID

	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		result.Code = http.StatusUnauthorized
		result.Message = "Token Required"
		ctx.AbortWithStatusJSON(result.Code, result)
		return
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	token, err := utils.ValidateUserJWTToToken(tokenString)

	if err != nil {
		if err.Error() == "invalid token" {
			result.Message = "Token Expired"
			ctx.AbortWithStatusJSON(http.StatusBadRequest, result)
		}
		result.Code = http.StatusInternalServerError
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, result)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		result.Message = "Invalid claims"
		ctx.AbortWithStatusJSON(result.Code, result)
	}
	messageProviderModel.User_id = claims["id"].(string)

	// Response data for successful registration
	Response := messageProviderModel

	r := q.MessageProviderRepositoryCommand.Updates(ctx, Response)
	if r.DB.Error != nil {
		// If there was an error, return Internal Server Error with error message
		result.Code = http.StatusInternalServerError
		ctx.AbortWithStatusJSON(result.Code, result)
		return
	}

	if r.DB.RowsAffected == 0 {
		// If there was an error, return Internal Server Error with error message
		result.Message = "Message Provider ID not available"
		ctx.AbortWithStatusJSON(result.Code, result)
		return
	}
	result = utils.ResultResponse{
		Code:    http.StatusOK,
		Data:    Response,
		Message: "Success Update MessageProvider",
		Status:  true,
	}
	// If messageprovider record was successfully saved, respond with messageprovider's registration data
	ctx.JSON(result.Code, result)

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
