package router

import (
	"github.com/gin-gonic/gin"
	"os"
)

func Initialize() {
	r := gin.Default()

	initializeRoutes(r)

	// Get the port from the environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := r.Run("0.0.0.0:" + port)
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:port
}
