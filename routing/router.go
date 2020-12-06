package routing

import "github.com/gin-gonic/gin"

// InitRouting initializes API routes
func InitRouting() {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "works",
		})
	})

	DefineUserRoutes(router)

	router.Run()
}
