#!/bin/bash

# Default values
EXTERNAL_REGISTRY="docker.io"
EXTERNAL_IMAGE="nginx:latest"

# Function to display usage
usage() {
    echo "Usage: $0 [--registry external.registry.com] [--image image:tag]"
    echo "  --registry  External registry (default: docker.io)"
    echo "  --image     Image to pull (default: nginx:latest)"
    echo "  --help      Show this help message"
    exit 1
}

# Parse long-form command-line options
while [[ $# -gt 0 ]]; do
    case "$1" in
        --registry)
            EXTERNAL_REGISTRY="$2"
            shift 2
            ;;
        --image)
            EXTERNAL_IMAGE="$2"
            shift 2
            ;;
        --help)
            usage
            ;;
        *)
            echo "Unknown option: $1"
            usage
            ;;
    esac
done

echo "Pulling image $EXTERNAL_IMAGE from $EXTERNAL_REGISTRY..."
docker pull "$EXTERNAL_REGISTRY/$EXTERNAL_IMAGE"

echo "Image pulled successfully!"
