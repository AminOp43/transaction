# Expense Tracker API

A RESTful API built with Go, Gin, PostgreSQL, JWT, and bcrypt for managing personal income and expense transactions.

Each user can create an account, log in, and manage only their own transactions.

## Features

* User signup and login
* Password hashing with bcrypt
* JWT authentication
* Create, read, update, and delete transactions
* User-specific transaction access
* Pagination for transaction lists
* PostgreSQL database
* Request timeout middleware
* Layered architecture

## Tech Stack

* Go
* Gin
* PostgreSQL
* JWT
* bcrypt
* golang-migrate

## Project Structure

```text
expense-tracker/
├── internal/
│   ├── domain/
│   ├── handler/
│   ├── middleware/
│   ├── repository/
│   │   └── postgres/
│   └── service/
├── migrations/
├── pkg/
│   └── db/
├── main.go
├── go.mod
└── README.md
```

## Architecture

The project follows a layered architecture:

```text
Handler → Service → Repository → PostgreSQL
```

* **Handler:** Handles HTTP requests and responses
* **Service:** Contains application and business logic
* **Repository:** Handles database operations
* **Middleware:** Handles authentication and request-related logic

## API Endpoints

### Authentication

```text
POST /user/signup
POST /user/login
```

### User

Authentication required:

```text
PUT    /user/
DELETE /user/
```

### Transactions

Authentication required:

```text
GET    /transaction/?page=1
GET    /transaction/:id
POST   /transaction/
PUT    /transaction/:id
DELETE /transaction/:id
```

## Authentication

After a successful login, the API returns a JWT token.

Protected routes require the token in the Authorization header:

```text
Authorization: Bearer <token>
```

## Transaction Example

```json
{
  "type": "expense",
  "amount": 250000,
  "category": "food",
  "description": "Dinner",
  "occurred_at": "2026-08-10T18:30:00Z"
}
```

The transaction type must be either:

```text
income
expense
```

## Pagination

Transactions are returned in pages.

Example:

```text
GET /transaction/?page=2
```

Each page currently contains 10 transactions.

## Environment Variables

Create a `.env` file in the project root:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_NAME=expense_tracker
DB_PASSWORD=your_password
DB_SSLMODE=disable

JWT_SECRET=your_jwt_secret
```

Do not commit the `.env` file to GitHub.

## Run the Project

Install dependencies:

```bash
go mod tidy
```

Run database migrations, then start the application:

```bash
go run .
```

The API runs on:

```text
http://localhost:8080
```

## Testing

Run all tests with:

```bash
go test ./...
```

## Security

* Passwords are hashed using bcrypt.
* Protected endpoints require JWT authentication.
* Users can access only their own transactions.
* JWT secrets and database credentials are stored in environment variables.

## Author

Developed as a Go backend learning project.
