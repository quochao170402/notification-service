package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/quochao170402/notification-service/internal/config"
	"github.com/quochao170402/notification-service/internal/service"
)

func main() {
	cfg := config.Load()

	mongoClient, err := service.NewMongoClient(
		cfg.MongoConfig.MongoURI,
		cfg.MongoConfig.MongoDB,
		cfg.MongoConfig.MongoUser,
		cfg.MongoConfig.MongoPass)

	if err != nil {
		log.Fatalf("Mongo connect failed: %v", err)
	}

	defer mongoClient.Disconnect(context.Background())

	fmt.Println("Connected to MongoDB:", cfg.MongoConfig)

	router := gin.Default()

	config.SetupRouters(router, mongoClient)

	log.Fatal(router.Run(":" + cfg.AppConfig.Port))
}
