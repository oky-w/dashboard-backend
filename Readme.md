# Go Dashboard Backend

## Overview

The Go Dashboard Backend is a REST API built to manage customer data, bank accounts, transactions, and other related information. This backend integrates with a Next.js frontend, enabling customers to view and interact with their data, including bank account details, pocket balances, and term deposits.

The backend is implemented using Go (Golang) and utilizes Gin (a lightweight web framework) for handling HTTP requests. It interfaces with a PostgreSQL database for managing customer and transaction data.

The architecture of the project follows the Hexagonal Architecture (also known as Ports and Adapters), which separates core business logic from external dependencies (such as databases and third-party services), ensuring scalability, maintainability, and testability.

Additionally, the project uses Docker for containerization, simplifying deployment and management in production environments. Postman is employed for API testing, while golangci-lint ensures code quality through static code analysis.

## Table of Contents

1. [Getting Started](#getting-started)
   - Installation
   - Running the Application
   - Setup Dummy Data
   - Running Tests
   - Running Linter (optional)
2. [How to Use Postman for API Testing](#how-to-run-postman-for-api-testing)
3. [Features](#features)
4. [System Design](#system-design)
5. [Tech Stack](#tech-stack)
6. [Folder Structure](#folder-structure)

## Getting Started

### Installation

1. **Clone the repository**

   - Clone the project to your local machine using the following command:

   ```bash
   https://github.com/okyws/dashboard-backend.git
   ```

2. **Install missing dependencies**

   - After cloning, navigate to the project directory and run the following command to tidy up the Go modules and install any missing dependencies:

   ```bash
   go mod tidy
   ```

3. **Set up environment variables**

   - Create a `.env` file in the project directory and set the required environment variables. The required environment variables can be copied from the `.env.example` file.

### Running the Application

- Start the application by running the following command in your terminal:

  ```bash
  go run main.go
  ```

### Setup Dummy Data

1. **Migrate the database** : To set up the database, run the following command:

```bash
go run main.go migrate
```

This command will migrate the database schema to the latest version.

2. **Seed Database** : To set up dummy data for testing and development, run the following command:

```bash
go run main.go seed
```

This command will seed the database with dummy data.

3. **Drop the database** : To drop the database, run the following command:

```bash
go run main.go drop
```

This command will drop the database schema.

### Running Tests

To test the application, you can use Postman (explained below) or run unit tests directly from the command line:

```bash
go test ./...
```

### Running Linter (optional)

To ensure code quality, the project uses golangci-lint for static analysis. To run the linter:

1. Ensure you have golangci-lint installed. You can install it using https://golangci-lint.run/welcome/install/
2. Navigate to the project directory.
3. Run the following command:
   ```bash
   golangci-lint run
   ```
   or
   ```bash
   golangci-lint run -c .golangci.yml
   ```

## How to Run Postman for API Testing

Postman is used for testing the API endpoints. Below is a step-by-step guide on how to set up and use Postman for testing the API.

### Step-by-Step Guide

### Step 1: Import Postman Collection

1. Open **Postman** and navigate to **File → Import** or click on the **Import** button.
2. Select the `Go-Dashboard.postman_collection.json` file from this repository and import it into Postman.

### Step 2: Create a Customer Record

1. In Postman, locate the **Customer** folder in the imported collection.
2. Select the **Create** request.
3. Fill in the required fields in the request body, such as `email`, `username`, `password`, and `role`.
4. Send the request.
5. If the data is valid, the API will respond with a success message and the created customer data. Otherwise, it will return an error message.

### Step 3: Create a Bank Account Record

1. Under the **Bank Account** folder, find the **Create** request.
2. Fill in the required fields in the request body, including the `user_id` (use the **customer ID** from the previous step), `account_type` (choose from `rekening-utama`, `saku`, `celengan`, `deposito`), and other necessary details.
3. Send the request.
4. If the request is successful, the new bank account data will be returned.

### Step 4: Create a Transaction Record

1. Locate the **Transaction** folder in Postman and select the **Create** request.
2. Fill in the required fields such as `from_account_number`, `to_account_number`, `transaction_type` (`transfer`, `deposit`, or `withdraw`), and transaction amount.
3. Send the request to create a transaction.
4. If the transaction is valid, you will get a success response.

## Features

The Go Dashboard Backend includes the following key features:

- **User Management**: Allows for CRUD operations (Create, Read, Update, Delete) for managing user data.
- **Customer Data Management**: Supports CRUD operations for customer information.
- **Bank Account Management**: Enables CRUD operations to manage bank account records.
- **Transaction Management**: Tracks transactions, including transfers, deposits, and withdrawals.
- **Pocket Information**: Handles pocket balances and related transactions.

## System Design

### 1. Hexagonal Architecture

The project follows the Hexagonal Architecture (also known as the Ports and Adapters pattern), which separates the core business logic from external systems, such as databases or third-party services. This architecture improves maintainability, testability, and scalability by ensuring that the core logic does not depend directly on external systems.

#### Key Components:

- **Adapters**: Handle communication between the core business logic and external systems. For example, HTTP handlers (Gin) and repositories (PostgreSQL).
- **Ports**: Define interfaces that the core business logic uses to interact with external systems. For example, repositories for database interaction, and services for interacting with third-party APIs.
- **Core Business Logic**: This contains the domain entities and business services responsible for implementing the application's business rules.

### 2. Folder Structure

The project is organized into the following structure, reflecting the Hexagonal Architecture principles:

```
dashboard-backend/
│── constants/                    # Common constants for the application
│── domain/                       # Domain entities and interfaces
│── dto/                          # Data transfer objects
│── adapter/handler/              # HTTP handlers
│── adapter/repository/           # Data access layer for interacting with the database
│── middleware/                   # Middleware for HTTP request logging, etc.
│── service/                      # Business logic services
│── log/                          # Logging output for the application
│── config/                       # Configuration files
│── routes/                       # HTTP route definitions
├── ports/                        # Interfaces for repositories and services
│── main.go                       # Main function as a loader
│── utils/                        # Utility functions
│── test/                         # Unit and integration tests
│── README.md                     # Project documentation
```

### 3. Dockerization

The project is containerized using Docker for easier deployment and scalability in production environments.

To build and run the Docker container, follow these steps:

1. **Build and Run the Docker container:**

   ```bash
   docker compose up --build -d
   ```

2. **Access the application:**

   - Open your web browser and navigate to `http://localhost:8080`.

3. **Stop and remove the Docker container:**

   ```bash
   docker compose down
   ```

## Tech Stack

### Backend

- **Language**: Go (Golang)
- **Framework**: Gin (HTTP web framework)

### Database

- **Database**: PostgreSQL for storing user, customer, bank account, and transaction data. SQLite for testing purposes
- **ORM**: GORM (Go ORM for interacting with PostgreSQL)

### Core Libraries and Dependencies

- **UUID**: For generating unique identifiers
- **Configuration Management**: For managing environment variables
- **Logging**: Structured logging
- **Encryption**: For secure password hashing
- **JWT**: For handling JSON Web Tokens for authentication
- **Cobra**: For command-line interface (CLI) support commands

### Testing and Quality Assurance

- **Unit Testing**: For unit tests and mocks
- **Static Code Analysis**: For linting and ensuring code quality

### Containerization and Deployment

- **Docker**: Containerization for consistent deployment across environments

### API Testing

- **Postman**: Used for API testing and interaction

## Expected Output

The expected output is a **REST API** that allows users to manage customer data, bank accounts, transactions, and other related information and it will be consumed by a **Next.js** frontend.

## Result / Summary Project

### Completed Features:

- **User Management**: Provides CRUD operations for user data.
- **Customer Data Management**: Supports CRUD operations for managing customer information.
- **Bank Account Management**: Manages bank account records with CRUD operations.
- **Transaction Management**: Tracks and manages customer transactions.
- **Pocket Information**: Handles pocket balance and related transactions.
