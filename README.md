﻿﻿**This is the backend code for GoRide in GoLang**

## Setup

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/SanchitDang/go_ride_backend-GoLang.git
   cd go_ride_backend
   ```

2. **Run the Application:**
   ```bash
   go run main.go
   ```

   The server will start listening on `http://localhost:8080`.

3. **Test Endpoints:**
   - User POST data: Send a POST request to `http://localhost:8080/post-user` with JSON data in the request body.
   - User GET data: Send a GET request to `http://localhost:8080/get-all-users` to retrieve all stored users.

## API Endpoints

### 1. POST Data (`/post-user`)

#### User Request:

```json
{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "phone": "123-456-7890"
}
```

#### Response:

```json
{
  "message": "Data received and stored successfully"
}
```

### 2. GET Data (`/get-all-users`)

#### User Response:

```json 
[
  {
    "name": "John Doe",
    "email": "john.doe@example.com",
    "phone": "123-456-7890",
  },
  {
    "name": "Jane Doe",
    "email": "jane.doe@example.com",
    "phone": "321-456-7890",
  },
]
```

## Dependencies

- [Gin](https://github.com/gin-gonic/gin): A powerful URL router and dispatcher for Go.
- [godotenv](https://github.com/joho/godotenv): A Go port of Ruby's dotenv library (Loads environment variables from .env file).
- [Viper](https://github.com/spf13/viper): Go configuration with fangs.
- [Gorm Postgres Driver](https://github.com/go-gorm/postgres): The PostgreSQL driver for GORM.
- [Gorm](https://github.com/go-gorm/gorm): The fantastic ORM library for Golang, aims to be developer-friendly.
