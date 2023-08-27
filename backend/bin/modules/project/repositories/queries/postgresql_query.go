package queries

import (
	"login-api-jwt/bin/modules/project"
	"login-api-jwt/bin/modules/project/models"
	"login-api-jwt/bin/pkg/databases"
	"login-api-jwt/bin/pkg/utils"

	"github.com/gin-gonic/gin"
)

// QueryRepository implements project.RepositoryQuery interface
type QueryRepository struct {
	ORM *databases.ORM
}

// NewQueryRepository creates a new instance of QueryRepository
func NewQueryRepository(orm *databases.ORM) project.RepositoryQuery {
	return &QueryRepository{
		ORM: orm,
	}
}

// FindOneByID retrieves a project record from database by ID

func (q QueryRepository) FindAll(ctx *gin.Context) utils.Result {
	var projectsModel []models.Project

	// Use ORM to find a project record by ID
	r := q.ORM.DB.Find(&projectsModel)

	// Prepare the result, including retrieved project data and database operation result
	output := utils.Result{
		Data: projectsModel,
		DB:   r,
	}
	return output

}
func (q QueryRepository) FindOneByID(ctx *gin.Context, id string) utils.Result {
	var projectModel models.Project

	// Use ORM to find a project record by ID
	r := q.ORM.DB.First(&projectModel, "id = ?", id)
	// Prepare the result, including retrieved project data and database operation result
	output := utils.Result{
		Data: projectModel,
		DB:   r,
	}
	return output

}

// // FindOneByName retrieves a project record from database by name
// func (q QueryRepository) FindOneByName(ctx *gin.Context, name string) utils.Result {
// 	var projectModel models.Project

// 	// Use ORM to find a project record by name
// 	r := q.ORM.DB.First(&projectModel, "name = ?", name)
// 	// Prepare the result, including retrieved project data and database operation result
// 	output := utils.Result{
// 		Data: projectModel,
// 		DB:   r,
// 	}
// 	return output

// }
