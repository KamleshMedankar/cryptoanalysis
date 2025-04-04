package workers

import (
	"crypto_analysis/config"
	"crypto_analysis/models"
	"encoding/json"
	"log"
	"math/rand"
	"time"
)

var JobQueue = make(chan models.Job, 100)

func NewJob(requestID string, request models.AnalysisRequest) models.Job {
	return models.Job{RequestID: requestID, Request: request}
}

func StartWorkerPool(workerCount int) {
	for i := 0; i < workerCount; i++ {
		go worker(i)
	}
}

func worker(workerID int) {
	for job := range JobQueue {
		log.Printf("[Worker %d] Processing job for %s", workerID, job.Request.WalletAddress)

		updateStatus(job.RequestID, "processing", nil)
		time.Sleep(2 * time.Second)

		if rand.Intn(10) < 1 {
			updateStatus(job.RequestID, "failed", nil)
			log.Printf("[Worker %d] Job failed for %s", workerID, job.Request.WalletAddress)
			continue
		}

		score := rand.Intn(100)
		updateStatus(job.RequestID, "completed", &score)
		log.Printf("[Worker %d] Risk Score: %d", workerID, score)
	}

}

func updateStatus(requestID, status string, score *int) {
	val, err := config.RedisClient.Get(config.Ctx, "request:"+requestID).Result()
	if err != nil {
		log.Printf("Error fetching job %s from Redis: %v", requestID, err)
		return
	}

	var data map[string]interface{}
	json.Unmarshal([]byte(val), &data)
	data["status"] = status
	if score != nil {
		data["risk_score"] = *score
	}

	newData, _ := json.Marshal(data)
	config.RedisClient.Set(config.Ctx, "request:"+requestID, newData, 0)
}
