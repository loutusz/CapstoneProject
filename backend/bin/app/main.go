package main

import (
	"log"
	userHandler "login-api-jwt/bin/modules/user/handlers"
	userRepositoryCommands "login-api-jwt/bin/modules/user/repositories/commands"
	userRepositoryQueries "login-api-jwt/bin/modules/user/repositories/queries"
	userUsecases "login-api-jwt/bin/modules/user/usecases"

	projectHandler "login-api-jwt/bin/modules/project/handlers"
	projectRepositoryCommands "login-api-jwt/bin/modules/project/repositories/commands"
	projectRepositoryQueries "login-api-jwt/bin/modules/project/repositories/queries"
	projectUsecases "login-api-jwt/bin/modules/project/usecases"

	"login-api-jwt/bin/pkg/databases"
	"login-api-jwt/bin/pkg/servers"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	// Create instance of Gin server
	srv := &servers.GinServer{}
	orm := &databases.ORM{}
	// userHandler := &handlers.UserServer{}

	// Get dsn from environment variables
	dsn := os.Getenv("DB_DSN")

	// Initialize Gin server
	srv.InitGin()
	orm.InitDB(dsn)
	// srv.InitTryRoutes()

	// Set up HTTP routes(orm) and handlers(srv)
	setUserHTTP(orm, srv)
	setProjectHTTP(orm, srv)

	// Start Gin server, listening on port 8050
	srv.Start(":8050", orm)
}

func setUserHTTP(orm *databases.ORM, srv *servers.GinServer) {
	// Create a user query repository and use case for reading user data
	userQueryRepository := userRepositoryQueries.NewQueryRepository(orm)
	userQueryUsecase := userUsecases.NewQueryUsecase(userQueryRepository, orm)

	// Create a user command repository and use case for writing user data
	userCommandRepository := userRepositoryCommands.NewCommandRepository(orm)
	userCommandUsecase := userUsecases.NewCommandUsecase(userCommandRepository, orm)

	// Initialize user HTTP handlers with query and command use cases, and link them with the Gin server
	userHandler.InitUserHTTPHandler(userQueryUsecase, userCommandUsecase, srv)
}

func setProjectHTTP(orm *databases.ORM, srv *servers.GinServer) {
	// Create a user query repository and use case for reading project data
	projectQueryRepository := projectRepositoryQueries.NewQueryRepository(orm)
	projectQueryUsecase := projectUsecases.NewQueryUsecase(projectQueryRepository, orm)

	// Create a project command repository and use case for writing project data
	projectCommandRepository := projectRepositoryCommands.NewCommandRepository(orm)
	projectCommandUsecase := projectUsecases.NewCommandUsecase(projectCommandRepository, orm)

	// Initialize project HTTP handlers with query and command use cases, and link them with the Gin server
	projectHandler.InitProjectHTTPHandler(projectQueryUsecase, projectCommandUsecase, srv)
}
