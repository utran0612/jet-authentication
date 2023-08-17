# jwt-authentication
This is a Golang-based authentication service using the gin-gonic framework and JWT (JSON Web Tokens).

## Installation

Before you begin, make sure you have Go installed on your machine. You can download and install it from the official Go website: [https://golang.org/dl/](https://golang.org/dl/)

### Running the Project

To run the authentication service, follow these steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/utran0612/jwt-authentication.git
   ```
2. Run go mod tidy to clean and update your module dependencies:
   ```bash
   go mod tidy
   ```
3. Build and run the application:
   ```bash
   go run main.go
   ```
4. The authentication service should now be running on http://localhost:8080.

### Usage

You can use this authentication service to manage user authentication and JWT generation. Make API requests to endpoints such as /signup, /login, /getUser and /getUsers (example endpoints).

### Contributing

Contributions are welcome! Feel free to fork this repository, make changes, and submit pull requests.

### License

This project is licensed under the MIT License - see the LICENSE file for details.

