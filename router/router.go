package router

import "github.com/gin-gonic/gin"

func Initialize() {
	r := gin.Default()

	initializeRoutes(r)

	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}
