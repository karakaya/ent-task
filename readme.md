## ent (company name masked) transaction processing api task



A transaction processing api that reliably handles user transactions by ensuring each transaction is processed only once. It uses PostgreSQL database transactions to prevent duplicate processing, reduce inconsistency, and improve reliability.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Setup](#Setup)
- [Usage](#usage)
- [Configuration](#configuration)
- [Testing](#testing)



## Prerequisites

Before you begin, ensure you have met the following requirements:

- **Docker**: [Download and install Docker](https://www.docker.com/get-started)
- **Docker Compose**: [Install Docker Compose](https://docs.docker.com/compose/install/)
- **Go 1.23.3**: [Install Go](https://golang.org/dl/) (sorry for allocating space for the new image :) )

## Setup & build & test

1. **Build**
   ```bash
   docker compose up -d
2. It will create the api container along with a postgres container.
3. Migrations will run automatically

4. **Testing** 
#####
Used mockery for testing. 
Use from the project directory.
```bash
   go test -v ./cmd/api-service/user/internal/...
   go test -v ./pkg/core/...
