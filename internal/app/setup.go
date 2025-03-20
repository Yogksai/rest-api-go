package app

import "github.com/gin-gonic/gin"

func setupRouters(r *gin.Engine) {
	apiGroup := r.Group("/api")
	v1Group := apiGroup.Group("/v1")
	routes.RegisterUserRoutes(v1Group)
}
