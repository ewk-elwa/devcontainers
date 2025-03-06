# Pulling from external registry and pushing to local registry
```bash
Step 1: Pull the Image
./pull-image-from-external.sh --registry docker.io --image nginx:latest

# Step 2: Push Image to Local K3s Registry
./push-image-to-local-registry.sh --local-registry localhost:5000 --image nginx:local

# Step 3: Deploy the Hello World App
./deploy-nginx-helloworld.sh --namespace default --deployment nginx-helloworld

# Verification
# Check local registry
curl -u admin:strongpassword http://localhost:5000/v2/_catalog

# Check deployment
kubectl get pods
kubectl get deployments
kubectl get services
```
