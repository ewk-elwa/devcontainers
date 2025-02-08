# Todo List Application

This is a simple Todo List application built with Go. It allows users to create, read, update, and delete todo items through a web interface.

## Project Structure

```
todo
├── src
│   ├── main.go          # Entry point of the application
│   ├── handlers         # Contains HTTP request handlers
│   │   └── todo.go
│   ├── models           # Defines the Todo model
│   │   └── todo.go
│   └── routes           # Sets up application routes
│       └── routes.go
├── Dockerfile           # Docker configuration for the application
├── deployment.yaml      # Kubernetes deployment configuration
├── go.mod               # Go module definition
└── README.md            # Project documentation
```

## Getting Started

### Prerequisites

- Go 1.16 or later
- Docker
- Kubernetes (optional for deployment)

### Building the Application

1. Clone the repository:
   ```
   git clone <repository-url>
   cd todo
   ```

2. Build the application:
   ```
   go build -o todo ./src/main.go
   ```

### Running the Application

To run the application locally, use the following command:
```
go run ./src/main.go
```
The server will start on `http://localhost:8080`.

### Docker

To build the Docker image, run:
```
docker build -t todo .
```

To run the Docker container:
```
docker run -p 8080:8080 todo
```

### Kubernetes Deployment

To deploy the application on Kubernetes, apply the deployment configuration:
```
kubectl apply -f deployment.yaml
```

### API Endpoints

- `POST /todos` - Create a new todo
- `GET /todos` - Retrieve all todos
- `GET /todos/{id}` - Retrieve a specific todo by ID
- `PUT /todos/{id}` - Update a specific todo by ID
- `DELETE /todos/{id}` - Delete a specific todo by ID

## License

This project is licensed under the MIT License.