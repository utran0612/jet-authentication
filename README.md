# jwt-authentication
This is a Golang-based authentication service using the gin-gonic framework and JWT (JSON Web Tokens).

## Installation

Before you begin, make sure you have Go installed on your machine. You can download and install it from the official Go website: [https://golang.org/dl/](https://golang.org/dl/)

### Installing Dependencies

To install the required dependencies for this project, follow these steps:

1. Install the gin-gonic framework:
   ```bash
   go get -u github.com/gin-gonic/gin
   ```
2. Run go mod tidy to clean and update your module dependencies:
   ```bash
   go mod tidy
   ```
   
### Running the Project

To run the authentication service, follow these steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/auth-service.git
   cd auth-service
   ```
2. Build and run the application:
   ```bash
   go run main.go
   ```
3. The authentication service should now be running on http://localhost:8080.

### Usage

You can use this authentication service to manage user authentication and JWT generation. Make API requests to endpoints such as /signup, /login, and /protected (example endpoints).

### Contributing

Contributions are welcome! Feel free to fork this repository, make changes, and submit pull requests.

### License

This project is licensed under the MIT License - see the LICENSE file for details.

