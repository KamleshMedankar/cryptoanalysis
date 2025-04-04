package controllers

import (
	"crypto_analysis/config"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckStatus(c *gin.Context) {
	requestID := c.Query("request_id")
	if requestID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing request_id"})
		return
	}

	key := "request:" + requestID
	val, err := config.RedisClient.Get(config.Ctx, key).Result()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Request not found"})
		return
	}

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding request data"})
		return
	}

	c.JSON(http.StatusOK, data)
}
