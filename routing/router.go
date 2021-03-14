package routing

import (
	"log"

	"github.com/gin-gonic/gin"
)

// InitRouting initializes API routes
func InitRouting() {
	router := gin.Default()

	v1 := router.Group("/api")

	v1.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	DefineAuthRoutes(v1.Group("/auth"))
	DefineUserRoutes(v1.Group("/users"))

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
