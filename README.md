# Hotel Reservation API

This a demonstartion of a RESTful API for hotel room booking and management system built with Go, Fiber, and MongoDB.

## 📋 Features

_Legend_: ✅ Implemented | 🚧 In Progress | ❌ Not Started

- ✅ User authentication and authorization using JWT
- ✅ Hotel management (CRUD operations)
- ✅ Room management and booking system
- 🚧 Admin dashboard (partially implemented)
- ❌ Database seeding and migration scripts
  [➡️ Jump to Current Progress section](#-current-progress)

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

## 🚧 Current Progress

The project is still in development. Here's what's implemented and how you can test it:

### Implemented Features

- **User Management**: Full CRUD operations for users
- **Hotel Management**: Create, read, and delete hotels
- **Room Management**: View and book available rooms
- **Booking System**: Create and cancel bookings

### Testing the API

You can test the API using tools like `curl` or Postman. Here are some example endpoints:

1. **User Registration**

   ```bash
   curl -X POST http://localhost:3000/api/v1/user -H "Content-Type: application/json" -d '{"email":"test@example.com", "password":"password123"}'
   ```

2. **User Authentication**

   ```bash
   curl -X POST http://localhost:3000/api/auth -H "Content-Type: application/json" -d '{"email":"test@example.com", "password":"password123"}'
   ```

   (Use the returned JWT token in the `Authorization: Bearer <token>` header for subsequent requests)

3. **View Available Hotels**

   ```bash
   curl http://localhost:3000/api/v1/hotel
   ```

4. **Book a Room**
   ```bash
   curl -X POST http://localhost:3000/api/v1/room/room-id/book -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d '{"from":"2025-08-01", "to":"2025-08-05"}'
   ```
