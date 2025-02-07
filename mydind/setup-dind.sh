#!/bin/bash

IMAGE_NAME="kdind:001"

# Check if the image already exists
if [[ "$(docker images -q $IMAGE_NAME 2> /dev/null)" == "" ]]; then
    echo "Image $IMAGE_NAME not found. Building the image..."
    docker build -t $IMAGE_NAME -f Dockerfile.dind .
else
    echo "Image $IMAGE_NAME already exists."
fi

# Run the Docker in Docker container if not already running
if [[ "$(docker ps -q -f name=mydind)" == "" ]]; then
    echo "Container mydind not running. Starting the container..."
    docker run --privileged -d --name mydind $IMAGE_NAME
else
    echo "Container mydind is already running."
fi

# Check if the k3d cluster already exists
if ! docker exec mydind k3d cluster list mypoc &> /dev/null; then
    echo "k3d cluster mypoc not found. Creating the cluster..."
    docker exec mydind k3d cluster create mypoc
else
    echo "k3d cluster mypoc already exists."
fi
