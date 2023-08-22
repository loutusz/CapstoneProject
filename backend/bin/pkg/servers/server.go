package servers

import (
	"errors"
	"github.com/gin-contrib/cors"
	"fmt"
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

func (s *GinServer) Start(port string, db *databases.ORM) error {
	if !s.Ready() && !db.Ready() {
		return errors.New("server isn't ready - make sure to init db and gin")
	}

	// Setup the endpoint
	endpoint := "0.0.0.0:" + port

	// s.Gin.Use(cors.Default())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "https://jico-api.up.railway.app/"} // Replace with your frontend origin
	config.AllowMethods = []string{"PUT","POST","GET","DELETE", "PATCH"}
	config.AllowHeaders = []string{"Origin"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true

	s.Gin.Use(cors.New(config))

	// Create a message indicating server startup
	startupMessage := fmt.Sprintf("Server is running on %s", endpoint)
	log.Println(startupMessage) // Log the startup message
	if err := http.ListenAndServe(endpoint, s.Gin.Handler()); err != nil {
		log.Fatal(err)
	}
	return nil
}
