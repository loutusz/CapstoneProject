package servers

import (
	"errors"
	"log"
	"login-api-jwt/bin/pkg/databases"
	"net/http"

	"github.com/gin-gonic/gin"
)

// make new gin server struct
type GinServer struct {
	Gin *gin.Engine
}

// initiate new gin server
func (s *GinServer) InitGin() *GinServer {
	g := gin.Default() // make new gin instance
	s.Gin = g          // assign new gin instace from g variable
	return s           // return new gin server
}

// ready
func (s *GinServer) Ready() bool {
	return s.Gin != nil // if GinServer instance have no engine, return false
}

func (s *GinServer) Start(ep string, db *databases.ORM) error {
	if !s.Ready() && !db.Ready() {
		return errors.New("server isn't ready - make sure to init db and gin")
	}

	if err := http.ListenAndServe(ep, s.Gin.Handler()); err != nil {
		log.Fatal(err)
	}
	return nil
}
