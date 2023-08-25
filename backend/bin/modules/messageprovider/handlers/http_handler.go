package handlers

import (
	"login-api-jwt/bin/modules/messageprovider"
	"login-api-jwt/bin/pkg/servers"
)

type MessageProviderHttpHandler struct {
	MessageProviderUsecaseQuery   messageprovider.UsecaseQuery
	MessageProviderUsecaseCommand messageprovider.UsecaseCommand
}

func InitMessageProviderHTTPHandler(uq messageprovider.UsecaseQuery, uc messageprovider.UsecaseCommand, s *servers.GinServer) {
	// Create an instance of MessageProviderHttpHandler with provided use cases
	handler := &MessageProviderHttpHandler{
		MessageProviderUsecaseQuery:   uq,
		MessageProviderUsecaseCommand: uc,
	}

	s.Gin.GET("/message-provider/", handler.MessageProviderUsecaseQuery.GetAccess)
	s.Gin.GET("/message-provider/all", handler.MessageProviderUsecaseQuery.GetAll)
	s.Gin.GET("/message-provider/id/:id", handler.MessageProviderUsecaseQuery.GetByID)
	s.Gin.POST("/message-provider/new", handler.MessageProviderUsecaseCommand.PostMessageProvider)
	s.Gin.PUT("/message-provider/edit/:id", handler.MessageProviderUsecaseCommand.PostMessageProvider)
	s.Gin.DELETE("/message-provider/id/:id", handler.MessageProviderUsecaseCommand.PostMessageProvider)
}
