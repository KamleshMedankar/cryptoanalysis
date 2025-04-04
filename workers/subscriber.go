// subscriber.go
package workers

import (
	"crypto_analysis/config"
	"crypto_analysis/models"
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)



func StartNATSConsumer() {
	_, err := config.NatsConn.Subscribe("crypto_analysis.jobs", func(msg *nats.Msg) {
		var payload models.JobPayload
		err := json.Unmarshal(msg.Data, &payload)
		if err != nil {
			log.Println("[NATS] Failed to parse job payload:", err)
			return
		}

		log.Println("[NATS] Received job for:", payload.Request.WalletAddress)

		job := NewJob(payload.RequestID, payload.Request)
		JobQueue <- job
	})

	if err != nil {
		log.Fatal("[NATS] Subscription error:", err)
	}

	log.Println("[NATS] Subscribed to crypto_analysis.jobs")
}
