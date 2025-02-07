#!/bin/bash

# Define variables
IMAGE_NAME="movie-svc"
DOCKERFILE_PATH="Dockerfile.movie-svc"
EXPOSED_PORT=8080

# Build the Docker image
cd movie-svc
docker build -t $IMAGE_NAME -f $DOCKERFILE_PATH .

# Run the Docker container and expose the web service port
docker run -d -p $EXPOSED_PORT:$EXPOSED_PORT $IMAGE_NAME
