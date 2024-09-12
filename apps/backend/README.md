# Bored API

This project is a backend service that uses Go, PostgreSQL, Redis, and Docker. It provides a set of API routes for user authentication, event management, chat functionality, and more. This documentation will help you set up the project and access various endpoints, including Swagger for API documentation and pgAdmin for database management.

## Prerequisites

- Docker and Docker Compose installed
- Makefile to simplify Docker commands

## Getting Started

1. **Clone the repository**:

   ```bash
   git clone https://github.com/montekkundan/bored.git
   cd bored/apps/backend
   ```

2. **Start the application**:

   Use the Makefile to build and start the containers.

   ```bash
   make start
   ```
   This will start all the necessary services (API, PostgreSQL, Redis, pgAdmin).

3. **Stop the application**:

   To stop the services:

   ```bash
   make stop
    ```
## API Endpoints

### Authentication
- **Login**: `POST /api/auth/login`
- **Register**: `POST /api/auth/register`
- **Verify Email**: `POST /api/auth/verify-email`
- **Verify Phone**: `POST /api/auth/verify-phone`
- **Enable Two-Factor Authentication (2FA)**: `POST /api/auth/enable-2fa`
- **Logout**: `POST /api/auth/logout`
- **Rotate Refresh Token**: `POST /api/auth/rotate-token`

### Users
- **Get All Users**: `GET /api/users/get-all`
- **Update User**: `PUT /api/users/update-user`

### Events
- **Get Events**: `GET /api/event`

### Tickets
- **Get Tickets**: `GET /api/ticket`

### Chats
- **Get Chats**: `GET /api/chat`

## Accessing the Swagger API Documentation

Once the application is running, you can view the Swagger UI for all the API routes.

- **Swagger UI**: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

**Swagger UI**:
Fiber: Go web framework
Docker: Containerization for all services

## PgAdmin

You can access the pgAdmin dashboard to manage the PostgreSQL database.

- **PgAdmin**: [http://localhost:5050](http://localhost:5050)