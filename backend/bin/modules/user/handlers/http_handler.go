package handlers

import (
	"login-api-jwt/bin/modules/user"
	"login-api-jwt/bin/pkg/servers"
)

type UserHttpHandler struct {
	UserUsecaseQuery   user.UsecaseQuery
	UserUsecaseCommand user.UsecaseCommand
}

func InitUserHTTPHandler(uq user.UsecaseQuery, uc user.UsecaseCommand, s *servers.GinServer) {
	// Create an instance of UserHttpHandler with provided use cases
	handler := &UserHttpHandler{
		UserUsecaseQuery:   uq,
		UserUsecaseCommand: uc,
	}

	// Define and register various routes and their corresponding handlers
	// These routes are associated with different user-related operations
	s.Gin.GET("/user/id/:id", handler.UserUsecaseQuery.GetByID)
	s.Gin.GET("/user/", handler.UserUsecaseQuery.GetAccess)
	s.Gin.POST("/user/SignUp", handler.UserUsecaseCommand.PostRegister)
	s.Gin.GET("/user/username/:username", handler.UserUsecaseQuery.GetByUsername)
	s.Gin.POST("/user/SignIn", handler.UserUsecaseCommand.PostLogin)
	s.Gin.GET("/user/all", handler.UserUsecaseQuery.GetAll)
	s.Gin.DELETE("/user/id/:id", handler.UserUsecaseCommand.DeleteUser)
}
