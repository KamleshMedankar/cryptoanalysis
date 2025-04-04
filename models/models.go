// models/request.go
package models

type AnalysisRequest struct {
	WalletAddress string `json:"wallet_address" binding:"required"`
	WalletChain   string `json:"wallet_chain" binding:"required"`
}

type JobPayload struct {
	RequestID string
	Request   AnalysisRequest
}

type Job struct {
	RequestID string
	Request   AnalysisRequest
}