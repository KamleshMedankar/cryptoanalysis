package config

import (
	"log"

	"github.com/nats-io/nats.go"
)

var NatsConn *nats.Conn

func InitNATS() {
	var err error
	NatsConn, err = nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	log.Println("Connected to NATS")
}
