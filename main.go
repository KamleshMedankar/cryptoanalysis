package main

import (
	"crypto_analysis/config"
	"crypto_analysis/routes"
	"crypto_analysis/workers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Redis and NATS
	config.InitRedis()
	config.InitNATS()

	// Start worker pool
	workers.StartWorkerPool(2) // Max concurrent processing: 2
	workers.StartNATSConsumer()

	// Setup Gin router
	router := gin.Default()
	routes.SetupRoutes(router)

	log.Println("Server running on :8080")
	router.Run(":8080")
}
