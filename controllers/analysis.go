package controllers

import (
	"crypto_analysis/config"
	"crypto_analysis/models"
	"encoding/json"
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AnalyzeWallet(c *gin.Context) {
	var request models.AnalysisRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a unique request ID
	requestID := uuid.New().String()

	// Store initial job status in Redis
	jobData := map[string]interface{}{
		"status":         "queued",
		"wallet_address": request.WalletAddress,
		"wallet_chain":   request.WalletChain,
	}
	jsonData, _ := json.Marshal(jobData)
	config.RedisClient.Set(config.Ctx, "request:"+requestID, jsonData, 0)

	// Send job to NATS as a JSON payload
	payload := map[string]interface{}{
		"RequestID": requestID,
		"Request":   request,
	}
	payloadBytes, _ := json.Marshal(payload)

	if err := config.NatsConn.Publish("crypto_analysis.jobs", payloadBytes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to publish job to NATS"})
		return
	}
	log.Printf("[NATS] Published job for wallet: %s\n", request.WalletAddress)
	// Respond with request ID
	c.JSON(http.StatusAccepted, gin.H{
		"message":    "queued",
		"request_id": requestID,
	})
}
