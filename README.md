﻿**This is the backend code for GoRide in GoLang**

## Setup

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/SanchitDang/go_ride_backend.git
   cd go_ride_backend
   ```

2. **Run the Application:**
   ```bash
   go run main.go
   ```

   The server will start listening on `http://localhost:8080`.

3. **Test Endpoints:**
   - User POST data: Send a POST request to `http://localhost:8080/post-user` with JSON data in the request body.
   - User GET data: Send a GET request to `http://localhost:8080/get-user` to retrieve stored data.

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

### 2. GET Data (`/get-user`)

#### User Response:

```json
[
  {
    "name": "John Doe",
    "email": "john.doe@example.com",
    "phone": "123-456-7890",
  },
 
]
```

## Dependencies

- [Gin](https://github.com/gin-gonic/gin): A powerful URL router and dispatcher for Go.

- [go-sqlite3](https://github.com/mattn/go-sqlite3): SQLite3 driver for Go using database/sql.

- [go-Viper](https://github.com/spf13/viper): Go configuration with fangs.

