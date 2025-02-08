#!/bin/bash

# Function to build Docker image for recipebook
build_recipebook() {
  echo "Building Docker image for recipebook..."
  docker build -t recipebook:latest ./recipebook
  echo "Docker image for recipebook built successfully."
}

# Function to build Docker image for todo
build_todo() {
  echo "Building Docker image for todo..."
  docker build -t todo:latest ./todo
  echo "Docker image for todo built successfully."
}

# Run both build functions in parallel
build_recipebook
build_todo

# Wait for both builds to finish
# wait

echo "Both Docker images built successfully."