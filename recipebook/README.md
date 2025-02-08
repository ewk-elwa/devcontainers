# RecipeBook

RecipeBook is a simple web application built with Go that allows users to view a list of recipes, view details of a specific recipe, and create new recipes. This project is designed to demonstrate the use of Go for building RESTful web services.

## Features

- View all recipes
- View a single recipe
- Create a new recipe

## Project Structure

```
recipebook
├── src
│   ├── main.go                # Entry point of the application
│   ├── controllers            # Contains HTTP request handlers
│   │   └── recipeController.go
│   ├── models                 # Defines the data model
│   │   └── recipe.go
│   ├── routes                 # Defines application routes
│   │   └── routes.go
│   └── services               # Contains business logic
│       └── recipeService.go
├── Dockerfile                 # Docker configuration for the application
├── deployment.yaml            # Kubernetes deployment configuration
├── go.mod                    # Go module definition
└── README.md                  # Project documentation
```

## Getting Started

### Prerequisites

- Go (version 1.16 or later)
- Docker
- Kubernetes (for deployment)

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/recipebook.git
   cd recipebook
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

### Running the Application

To run the application locally, use the following command:

```
go run src/main.go
```

The server will start on `http://localhost:8080`.

### Docker

To build the Docker image, run:

```
docker build -t recipebook .
```

To run the Docker container, use:

```
docker run -p 8080:8080 recipebook
```

### Kubernetes Deployment

To deploy the application on Kubernetes, apply the deployment configuration:

```
kubectl apply -f deployment.yaml
```

## Usage

- **GET /recipes**: Retrieve a list of all recipes.
- **GET /recipes/{id}**: Retrieve details of a specific recipe by ID.
- **POST /recipes**: Create a new recipe. (Include JSON body with recipe details)

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for details.