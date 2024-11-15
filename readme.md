## ent (company name masked) transaction processing api task



A transaction processing api that reliably handles user transactions by ensuring each transaction is processed only once. It uses PostgreSQL database transactions to prevent duplicate processing, reduce inconsistency, and improve reliability.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Setup & Configuration](#setup)
- [Usage](#usage)
- [Testing](#testing)



## Prerequisites

Before you begin, ensure you have met the following requirements:

- **Docker**: [Download and install Docker](https://www.docker.com/get-started)
- **Docker Compose**: [Install Docker Compose](https://docs.docker.com/compose/install/)
- **Go 1.23.3**: [Install Go](https://golang.org/dl/) (sorry for allocating space for the new image :) )

## Setup
### Build & Test

1. **Build**
   ```bash
   docker compose up -d
2. It will create the api container along with a postgres container. Ports are configurable from the file docker-compose.yaml and exposed from :8080. 
3. Migrations will run automatically and insert users with id:1,2,3
4. Default users doesn't have any balance, will return 0

## Usage
1. **cURL examples for usage**

   Add Positive Balance 10 (user id 1)
   ```bash
   curl --location 'localhost:8080/user/1/transaction' --header 'Source-Type: payment' --header 'Content-Type: application/json' --data '{"state":"win","amount":"10.0","transactionId":"transaction-id-1"}'
      ```

   Get Balance (userId 1)
   ```bash
   curl --location 'localhost:8080/user/1/balance'
   ```
   Add Negative Balance 5 (user id 1)
   ```bash
   curl --location 'localhost:8080/user/1/transaction' --header 'Source-Type: payment' --header 'Content-Type: application/json' --data '{"state":"lose","amount":"5.0","transactionId":"transaction-id-1"}'
      ```

## Testing
1. **Testing** 
   #####
   Used mockery for testing. Use from the project directory.
   ```bash
   go test -v ./cmd/api-service/user/internal/...
   go test -v ./pkg/core/...
