
# AuthService

AuthService is a Go-based authentication service that provides user registration and login functionalities. It uses the Gin framework for handling HTTP requests and PostgreSQL for data storage.

## Project Structure

- `internal/api/controllers`: Contains the controllers for handling HTTP requests.
- `internal/api/middlewares`: Contains middleware for request processing.
- `internal/api/routes`: Contains route definitions.
- `internal/config`: Contains configuration management.
- `internal/pkg/auth`: Contains authentication logic.
- `internal/pkg/token`: Contains token management logic.
- `internal/pkg/users`: Contains user management logic.
- `internal/rlog`: Contains logging utilities.
- `internal/storage/postgres`: Contains PostgreSQL storage implementation.

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/viperh/authService.git
    cd authService
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Set up the configuration file:
    - Create .env file and update the configuration as needed.

4. Run the application:
    ```sh
    go run main.go
    ```

## Endpoints

### Authentication

#### Register

- **URL**: `/auth/register`
- **Method**: `POST`
- **Request Body**:
    ```json
    {
        "email": "user@example.com",
        "username": "username",
        "password": "password",
        "firstname": "First",
        "lastname": "Last"
    }
    ```
- **Response**:
    - **Success**: `201 Created`
        ```json
        {
            "token": {
                "access_token": "access_token_value",
                "refresh_token": "refresh_token_value"
            }
        }
        ```
    - **Error**: `400 Bad Request`
        ```json
        {
            "error": "Error message"
        }
        ```

#### Login

- **URL**: `/auth/login`
- **Method**: `POST`
- **Request Body**:
    ```json
    {
        "username": "username",
        "password": "password"
    }
    ```
- **Response**:
    - **Success**: `200 OK`
        ```json
        {
            "token": {
                "access_token": "access_token_value",
                "refresh_token": "refresh_token_value"
            }
        }
        ```
    - **Error**: `401 Unauthorized`
        ```json
        {
            "error": "Invalid credentials"
        }
        ```

## Configuration

Example .env File: 
```sh
DB_USER=example
DB_PASSWORD=example
DB_NAME=example
DB_HOST=example
DB_PORT=5432
DB_SSL_MODE=disable
LOG_LEVEL=dev # || prod
SERVER_PORT=3000
JWT_KEY=example
```

## Info

## Access token expiry: 15m
## Refresh token expiry: 7d
```
