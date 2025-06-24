# OMNIFUL-PROJECT

A microservices-based backend system for inventory and order management, consisting of two main services:
- **IMS**: Inventory Management Service
- **OMS**: Order Management Service

## Table of Contents
- [Project Overview](#project-overview)
- [Directory Structure](#directory-structure)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Setup & Installation](#setup--installation)
- [Running the Services](#running-the-services)
- [Database Migrations](#database-migrations)
- [API Overview](#api-overview)
- [Contributing](#contributing)
- [License](#license)

## Project Overview
This project provides a scalable backend for managing inventory, sellers, SKUs, hubs, tenants, and orders. It is split into two Go-based microservices:
- **ims/**: Handles inventory, sellers, SKUs, hubs, and tenants.
- **oms/**: Handles orders, CSV uploads, and webhooks.

## Directory Structure
```
OMNIFUL-PROJECT/
├── docker-compose.yml         # Multi-service orchestration
├── ims/                      # Inventory Management Service
│   ├── configs/              # Configuration files
│   ├── handlers/             # HTTP handlers
│   ├── main.go               # Service entrypoint
│   ├── migrations/           # SQL migrations
│   ├── models/               # Data models
│   ├── myContext/            # Context utilities
│   ├── myDb/                 # Database utilities
│   ├── routes/               # Route definitions
│   └── services/             # Business logic
├── oms/                      # Order Management Service
│   ├── client/               # External service clients (Kafka, S3, SQS)
│   ├── configs/              # Configuration files
│   ├── handlers/             # HTTP handlers
│   ├── main.go               # Service entrypoint
│   ├── model/                # Data models
│   ├── myContext/            # Context utilities
│   ├── myDB/                 # Database utilities
│   ├── routes/               # Route definitions
│   ├── services/             # Business logic
│   └── uploads/              # Uploaded files
└── README.md                 # Project documentation
```

## Features
- Multi-tenant inventory management
- Seller, SKU, and hub management
- Order processing and CSV uploads
- Webhook integration
- Kafka, S3, and SQS client support (OMS)
- Database migrations
- Docker Compose for local development

## Prerequisites
- [Go](https://golang.org/) 1.18+
- [Docker](https://www.docker.com/) & [Docker Compose](https://docs.docker.com/compose/)
- [PostgreSQL](https://www.postgresql.org/) (if not using Docker Compose)

## Setup & Installation
1. **Clone the repository:**
   ```sh
   git clone <repo-url>
   cd OMNIFUL-PROJECT
   ```
2. **Create required databases and cloud resources:**
   - **PostgreSQL:** Create a database named `ims_db` for the Inventory Management Service.
   - **MongoDB:** Create a database named `oms_db` for the Order Management Service.
   - **LocalStack:**
     - Create an S3 bucket (for file uploads, e.g., `oms-csv-bucket`).
     - Create an SQS queue (for order processing, e.g., `create-bulk-order`).
   > Ensure LocalStack is running locally and accessible. You can use the AWS CLI or LocalStack's web UI to create the bucket and queue.
3. **Copy and edit configuration files as needed.**
4. **Install Go dependencies:**
   ```sh
   cd ims && go mod download
   cd ../oms && go mod download
   cd ..
   ```

## Running the Services
### Using Docker Compose
1. Make sure Docker is running.
2. Start all services:
   ```sh
   docker-compose up --build
   ```

### Running Locally (without Docker)
1. Start your PostgreSQL instance and update configs as needed.
2. Run database migrations (see below).
3. Start each service:
   ```sh
   cd ims && go run main.go
   # In a new terminal:
   cd oms && go run main.go
   ```

## Database Migrations
- Migration files are in `ims/migrations/`.
- Use a migration tool like [golang-migrate](https://github.com/golang-migrate/migrate) to apply migrations:
  ```sh
  migrate -path ims/migrations -database <db-url> up
  ```

## API Overview
- **IMS**: Manages tenants, sellers, hubs, SKUs, and inventory.
- **OMS**: Manages orders, CSV uploads, and webhooks.
- See the `handlers/` and `routes/` directories in each service for endpoint details.

## Contributing
1. Fork the repository
2. Create a new branch (`git checkout -b feature/your-feature`)
3. Commit your changes
4. Push to your branch and open a Pull Request

## License
[MIT](LICENSE) (or specify your license here)