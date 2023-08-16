package usecases

import (
	"errors"
	"login-api-jwt/bin/modules/user"
	"login-api-jwt/bin/modules/user/models"
	"login-api-jwt/bin/pkg/databases"
	"login-api-jwt/bin/pkg/utils/validators"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CommandUsecase implements user.UsecaseCommand interface
type CommandUsecase struct {
	UserRepositoryCommand user.RepositoryCommand
	ORM                   *databases.ORM
}

// NewCommandUsecase creates a new instance of CommandUsecase
func NewCommandUsecase(q user.RepositoryCommand, orm *databases.ORM) user.UsecaseCommand {
	return &CommandUsecase{
		UserRepositoryCommand: q,
		ORM:                   orm,
	}
}

// PostRegister handles user registration
func (q CommandUsecase) PostRegister(ctx *gin.Context) {
	var userModel models.User
	err := ctx.ShouldBind(&userModel)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	// Generate a unique ID for user
	userModel.ID = uuid.NewString()

	// Capitalize first letter of user's name
	userModel.Name = strings.Title(userModel.Name)

	// Validate user's email format
	validEmail := validators.IsValidEmail(userModel.Email)
	validUsername := validators.IsValidUsername(userModel.Username)
	ValidPassword := validators.IsValidPassword(userModel.Password)

	if !validEmail {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "email not valid"})
		return
	}

	if !validUsername {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "username not valid"})
		return
	}

	if !ValidPassword {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "password not valid"})
		return
	}

	// Hash user's password before storing it in the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userModel.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	userModel.Password = string(hashedPassword)

	// Create user record in the database
	r := q.UserRepositoryCommand.Create(ctx, userModel)
	if r.DB.Error != nil {
		// Check if the error is due to a duplicate email or username
		if strings.Contains(r.DB.Error.Error(), "duplicate key value violates unique constraint \"users_email_key\"") {
			// If data is already found, abort with status "email or username already used"
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "email or username already used"})
			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, r.DB.Error)
		return
	}

	// Response data for successful registration
	userRegisterResponse := models.RegisterResponse{
		ID:       userModel.ID,
		Name:     userModel.Name,
		Username: userModel.Username,
		Email:    userModel.Email,
	}

	// Save user record again after successful registration
	r = q.UserRepositoryCommand.Save(ctx, userModel)

	// Check if an error occurred while saving
	if r.DB.Error != nil {
		// If there was an error, return Internal Server Error with error message
		ctx.AbortWithError(http.StatusInternalServerError, r.DB.Error)
		return
	}
	// If user record was successfully saved, respond with user's registration data
	ctx.JSON(http.StatusOK, userRegisterResponse)
}

// PostLogin handles user login
func (q CommandUsecase) PostLogin(ctx *gin.Context) {
	var userLoginRequest models.LoginRequest
	err := ctx.ShouldBind(&userLoginRequest)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Find user's password hash by username
	r := q.UserRepositoryCommand.FindPassword(ctx, userLoginRequest.Username)
	if r.DB.Error != nil {
		if errors.Is(r.DB.Error, gorm.ErrRecordNotFound) {
			// If data is not found in the database, abort with status Unauthorized
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		ctx.AbortWithError(http.StatusInternalServerError, r.Error)
	}

	// Compare the provided password with the hashed password in the database
	err = bcrypt.CompareHashAndPassword([]byte(r.Password), []byte(userLoginRequest.Password))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Incorrect username or password"})
		return
	}

	// Create a new JWT token for user
	token := jwt.New(jwt.SigningMethodHS256) // create new jwt token

	// Initialize claims variable as a map to hold JWT claims
	claims := token.Claims.(jwt.MapClaims)

	// Set claims in JWT token payload
	claims["id"] = r.Data.ID
	claims["username"] = r.Data.Username
	claims["name"] = r.Data.Name
	claims["email"] = r.Data.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Sign token with JWT secret key
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		// fmt.Println("found jwt error")
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Create a new instance of LoginResponse model, initializing its fields with data
	userLoginResponse := models.LoginResponse{
		ID:          r.Data.ID,
		Email:       r.Data.Email,
		Name:        r.Data.Name,
		Username:    r.Data.Username,
		AccessToken: t,
	}

	// Respond to request with an HTTP 200 OK status code and userLoginResponse data in JSON format
	ctx.JSON(http.StatusOK, userLoginResponse)
}
