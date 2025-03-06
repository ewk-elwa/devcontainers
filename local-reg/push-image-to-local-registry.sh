#!/bin/bash

# Default values
LOCAL_REGISTRY="localhost:5000"
LOCAL_IMAGE="nginx:local"
EXTERNAL_IMAGE="nginx:latest"

# Function to display usage
usage() {
    echo "Usage: $0 [--local-registry registry.local] [--image new_image:tag]"
    echo "  --local-registry  Local K3s registry (default: localhost:5000)"
    echo "  --image           New image tag (default: nginx:local)"
    echo "  --help            Show this help message"
    exit 1
}

# Parse long-form command-line options
while [[ $# -gt 0 ]]; do
    case "$1" in
        --local-registry)
            LOCAL_REGISTRY="$2"
            shift 2
            ;;
        --image)
            LOCAL_IMAGE="$2"
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

# Tag the image for local registry
echo "Tagging image as $LOCAL_REGISTRY/$LOCAL_IMAGE..."
docker tag "$EXTERNAL_IMAGE" "$LOCAL_REGISTRY/$LOCAL_IMAGE"

# Push to the local registry
echo "Pushing image to local registry..."
docker push "$LOCAL_REGISTRY/$LOCAL_IMAGE"

echo "Image successfully pushed to $LOCAL_REGISTRY/$LOCAL_IMAGE!"
