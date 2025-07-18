# Hotel Reservation API

This a demonstartion of a RESTful API for hotel room booking and management system built with Go, Fiber, and MongoDB.

## 📋 Features

- User authentication and authorization using JWT
- Hotel and room management (CRUD operations)
- Room booking and reservation system
- Admin dashboard for managing bookings
- Database seeding and migration scripts

## 🚀 Getting Started

### Prerequisites

- Go 1.16 or higher
- MongoDB 4.4 or higher (locally or in Docker)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/hotel-reservation-api.git
   cd hotel-reservation-api
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

3. Set up environment variables:
   ```bash
   cp .env.example .env
   ```
   Then edit `.env` with your configuration:
   ```
   HTTP_LISTEN_ADDRESS=:3000
   JWT_SECRET=your-secret-key-here
   MONGO_DB_NAME=hotel-reservation
   MONGO_DB_URL=mongodb://localhost:27017
   ```

### Running with Docker

1. Start MongoDB using Docker:

   ```bash
   docker run --name mongodb -d -p 27017:27017 mongo:latest
   ```

2. Build and run the application:
   ```bash
   docker build -t hotel-reservation-api .
   docker run -p 3000:3000 --link mongodb hotel-reservation-api
   ```

### Running Locally

1. Start the application:

   ```bash
   go run main.go
   ```

2. The API will be available at `http://localhost:3000`

### Makefile Commands

The project includes a Makefile building on top of go default commands:

- `make build`: Builds the application binary
- `make run`: Builds and runs the application
- `make test`: Runs all tests in the project

## 🛠️ Tech Stack

- **Framework**: [Fiber](https://gofiber.io/)
- **Database**: MongoDB
- **Authentication**: JWT
- **Containerization**: Docker
