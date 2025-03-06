#!/bin/bash

# Default values
NAMESPACE="default"
DEPLOYMENT_NAME="nginx-helloworld"
LOCAL_REGISTRY="localhost:5000"
IMAGE_NAME="nginx:local"

# Function to display usage
usage() {
    echo "Usage: $0 [--namespace namespace] [--deployment name] [--local-registry registry.local] [--image image:tag]"
    echo "  --namespace      Kubernetes namespace (default: default)"
    echo "  --deployment     Deployment name (default: nginx-helloworld)"
    echo "  --local-registry Local K3s registry (default: localhost:5000)"
    echo "  --image          Image name (default: nginx:local)"
    echo "  --help           Show this help message"
    exit 1
}

# Parse long-form command-line options
while [[ $# -gt 0 ]]; do
    case "$1" in
        --namespace)
            NAMESPACE="$2"
            shift 2
            ;;
        --deployment)
            DEPLOYMENT_NAME="$2"
            shift 2
            ;;
        --local-registry)
            LOCAL_REGISTRY="$2"
            shift 2
            ;;
        --image)
            IMAGE_NAME="$2"
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

# Create deployment YAML
cat <<EOF > nginx-deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: $DEPLOYMENT_NAME
  namespace: $NAMESPACE
spec:
  replicas: 1
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: $LOCAL_REGISTRY/$IMAGE_NAME
        imagePullPolicy: Always
        ports:
        - containerPort: 80
EOF

# Apply the deployment
echo "Deploying Nginx Hello World..."
kubectl apply -f nginx-deployment.yaml

echo "Deployment created successfully!"
