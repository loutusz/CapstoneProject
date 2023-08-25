package usecases

import (
	"errors"
	"fmt"
	"login-api-jwt/bin/modules/messageprovider"
	"login-api-jwt/bin/pkg/databases"
	"login-api-jwt/bin/pkg/utils"
	"math"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// QueryUsecase implements messageprovider.UsecaseQuery interface
type QueryUsecase struct {
	MessageProviderRepositoryQuery messageprovider.RepositoryQuery
	ORM                            *databases.ORM
}

// NewQueryUsecase creates a new instance of QueryUsecase
func NewQueryUsecase(q messageprovider.RepositoryQuery, orm *databases.ORM) messageprovider.UsecaseQuery {
	return &QueryUsecase{
		MessageProviderRepositoryQuery: q,
		ORM:                            orm,
	}
}

// GetByID retrieves messageprovider data by ID and responds with the result
func (q QueryUsecase) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")

	// Call FindOneByID method to retrieve messageprovider data by ID
	ret := q.MessageProviderRepositoryQuery.FindOneByID(ctx, id)
	// If there was an error during query, abort with a Bad Request status
	if ret.DB.Error != nil {
		if errors.Is(ret.DB.Error, gorm.ErrRecordNotFound) {
			// If data is not found in the database, abort with status Unauthorized
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("MessageProvider with id %s not found", id)})
			return
		}

		ctx.AbortWithError(http.StatusBadRequest, ret.DB.Error)
		return
	}

	// Respond with retrieved messageprovider data in JSON format
	res := ret.Data
	ctx.JSON(http.StatusOK, res)
}

// Get All retrieves All messageprovider data with pagination result
func (q QueryUsecase) GetAll(ctx *gin.Context) {
	var totalCount, page, limit int
	var err error

	page, err = strconv.Atoi(ctx.Query("page"))
	// handling when error set default page value 1
	if err != nil || page < 1 {
		page = 1
	}

	// handling when error set default limit value 10
	limit, err = strconv.Atoi(ctx.Query("limit"))
	if err != nil || limit == 0 {
		limit = 10
	}

	// count data messageprovider
	var count = q.MessageProviderRepositoryQuery.CountData(ctx)

	// parsing result count data to Integer
	if reflect.ValueOf(count.Data).CanInt() {
		totalCount = int(reflect.ValueOf(count.Data).Int())
	}

	// calculate total page
	var totalPage int = int(math.Ceil(float64(totalCount) / float64(limit)))
	// handling if max total page
	if page > totalPage {
		page = totalPage
	}

	// set skip for offset data
	var skip = limit * (page - 1)
	var result utils.ResultResponsePagination = utils.ResultResponsePagination{
		Code:      http.StatusBadRequest,
		Data:      nil,
		Limit:     limit,
		Page:      page,
		TotalData: totalCount,
		TotalPage: totalPage,
		Message:   "Failed Get Data MessageProvider",
		Status:    false,
	}

	if totalCount == 0 {
		result.Code = http.StatusNotFound
		result.Message = "Data Not Found"
		ctx.AbortWithStatusJSON(result.Code, result)
		return
	}

	getMessageProviderData := q.MessageProviderRepositoryQuery.FindAll(ctx, skip, limit)
	if getMessageProviderData.DB.Error != nil {
		ctx.AbortWithStatusJSON(result.Code, result)
		return
	}

	result = utils.ResultResponsePagination{
		Code:      http.StatusOK,
		Data:      getMessageProviderData.Data,
		Limit:     limit,
		Page:      page,
		TotalData: totalCount,
		TotalPage: totalPage,
		Message:   "Success Get Data MessageProvider",
		Status:    true,
	}
	ctx.JSON(http.StatusOK, result)
	return
}

// GetAccess responds with a success message indicating messageprovider access
func (q QueryUsecase) GetAccess(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "messageprovider access success"})
}

func (q QueryUsecase) GetUserOwned(ctx *gin.Context) {
	id := ctx.Param("id")
	var totalCount, page, limit int
	var err error

	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		return
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	_, err = utils.ValidateUserJWTToToken(tokenString)

	if err != nil {
		if err.Error() == "invalid token" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	page, err = strconv.Atoi(ctx.Query("page"))
	// handling when error set default page value 1
	if err != nil || page < 1 {
		page = 1
	}

	// handling when error set default limit value 10
	limit, err = strconv.Atoi(ctx.Query("limit"))
	if err != nil || limit == 0 {
		limit = 10
	}

	// count data messageprovider
	var count = q.MessageProviderRepositoryQuery.CountData(ctx)

	// parsing result count data to Integer
	if reflect.ValueOf(count.Data).CanInt() {
		totalCount = int(reflect.ValueOf(count.Data).Int())
	}

	// calculate total page
	var totalPage int = int(math.Ceil(float64(totalCount) / float64(limit)))
	// handling if max total page
	if page > totalPage {
		page = totalPage
	}

	// set skip for offset data
	var skip = limit * (page - 1)
	var result utils.ResultResponsePagination = utils.ResultResponsePagination{
		Code:      http.StatusBadRequest,
		Data:      nil,
		Limit:     limit,
		Page:      page,
		TotalData: totalCount,
		TotalPage: totalPage,
		Message:   "Failed Get Data MessageProvider",
		Status:    false,
	}

	if totalCount == 0 {
		result.Code = http.StatusNotFound
		result.Message = "Data Not Found"
		ctx.AbortWithStatusJSON(result.Code, result)
		return
	}

	getMessageProviderData := q.MessageProviderRepositoryQuery.FindByUserID(ctx, id, skip, limit)
	if getMessageProviderData.DB.Error != nil {
		ctx.AbortWithStatusJSON(result.Code, result)
		return
	}

	result = utils.ResultResponsePagination{
		Code:      http.StatusOK,
		Data:      getMessageProviderData.Data,
		Limit:     limit,
		Page:      page,
		TotalData: totalCount,
		TotalPage: totalPage,
		Message:   "Success Get Owned Data MessageProvider",
		Status:    true,
	}
	ctx.JSON(http.StatusOK, result)
	return
}
