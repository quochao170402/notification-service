package config

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quochao170402/notification-service/internal/api/handler"
	"github.com/quochao170402/notification-service/internal/repository"
	"github.com/quochao170402/notification-service/internal/service"
)

// NewRouter returns a configured Gin engine.
func NewRouter(mongoClient *service.Mongo) *gin.Engine {
	r := gin.Default()
	// Health check route
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	return r
}

func SetupRouters(r *gin.Engine, mongoClient *service.Mongo) {
	// Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(CORSMiddleware())

	// Create repositories
	repo := repository.NewNotificationRepository(mongoClient)

	// API v1 group
	v1 := r.Group("/api/v1")
	{
		notifications := v1.Group("/notifications")
		{
			handler.RegisterTaskRoutes(notifications, repo)
		}

	}
}

func CORSMiddleware() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
}
