package routing

import "github.com/gin-gonic/gin"

// InitRouting initializes API routes
func InitRouting() {
	router := gin.Default()

	v1 := router.Group("/api")

	v1.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "healthy",
		})
	})

	DefineAuthRoutes(v1.Group("/auth"))
	DefineUserRoutes(v1.Group("/users"))

	router.Run()
}
