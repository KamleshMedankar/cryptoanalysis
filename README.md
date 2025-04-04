# 🛡️ Crypto Wallet Risk Analysis Service

A microservice built in Golang using the Gin framework that analyzes the risk score of crypto wallets. It uses **Redis** for storing request data, **NATS** as a messaging queue, and **worker pools** to process jobs asynchronously.

---

## 🚀 Features
- Accepts wallet address and chain to analyze risk.
- Uses NATS to enqueue jobs for async processing.
- Worker pool to handle analysis tasks.
- Risk score calculation with job status tracking in Redis.

---

## ⚙️ Tech Stack
- 🐹 Golang (Gin Framework)
- 📨 NATS (Messaging Queue)
- 🟥 Redis (Key-value Store)
- 🔧 Worker Pool (Goroutines)

---

## 📦 Dependencies

Before running, make sure the following are installed:

- **Go** 1.18+
- **Redis** (Default port: `6379`)
- **NATS Server** (Default port: `4222`)

Install Go dependencies:

```bash
go mod tidy

🧪 Setup & Testing Instructions
1️⃣ Clone the Repo

git clone https://github.com/your-username/crypto-analysis.git
cd crypto-analysis

2️⃣ Start Redis and NATS
Make sure Redis and NATS servers are running:

# Start Redis
redis-server

# Start NATS
nats-server -p 4222

3️⃣ Run the Service
go run .

You should see:
Connected to Redis
Connected to NATS
[NATS] Subscribed to crypto_analysis.jobs
Server running on :8080

📮 API Documentation
🔍 Analyze Wallet
POST /analyze

Submit a wallet for analysis.

Request Body:
{
  "wallet_address": "0x123456",
  "wallet_chain": "Ethereum"
}
Response:
{
  "request_id": "uuid-string",
  "status": "queued"
}

📊 Check Status
GET /status?request_id=<uuid>

Fetch the status and risk score for a request.

Response (example):
{
  "status": "completed",
  "wallet_address": "0x123456",
  "wallet_chain": "Ethereum",
  "risk_score": 81
}

📁 Project Structure

.
├── config/         # Redis & NATS setup
├── controllers/    # Gin route handlers
├── models/         # Data models
├── routes/         # API routing
├── workers/        # NATS consumer and job workers
├── main.go         # App entry point

✅ Testing the Flow
1. Send a Wallet for Analysis:

curl -X POST http://localhost:8080/analyze \
-H "Content-Type: application/json" \
-d '{"wallet_address":"0x123456", "wallet_chain":"Dollor"}'

2. Check Redis for Status:
redis-cli
> get request:<request_id>

3. Check Status via API:
curl "http://localhost:8080/status?request_id=<request_id>"

🙋‍♂️ Author
👤 Kamlesh Krishna Medankar

