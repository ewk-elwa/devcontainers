#!/bin/bash

# Define image names and tar file names
MARIADB_IMAGE="mariadb:latest"
MOVIE_SVC_IMAGE="movie-svc:latest"
MARIADB_TAR="mariadb_image.tar"
MOVIE_SVC_TAR="movie_svc_image.tar"
CONTAINER_NAME="mydind"

# Save images to tar files
docker save -o $MARIADB_TAR $MARIADB_IMAGE
docker save -o $MOVIE_SVC_TAR $MOVIE_SVC_IMAGE

# Copy tar files into the mydind container
docker cp $MARIADB_TAR $CONTAINER_NAME:/root/
docker cp $MOVIE_SVC_TAR $CONTAINER_NAME:/root/

# Clean up local tar files
# rm $MARIADB_TAR
# rm $MOVIE_SVC_TAR

echo "Images saved and copied to the mydind container successfully."

# Load images from tar files inside the mydind container
docker exec $CONTAINER_NAME docker load -i /root/$MARIADB_TAR
docker exec $CONTAINER_NAME docker load -i /root/$MOVIE_SVC_TAR

# Clean up tar files inside the mydind container
# docker exec $CONTAINER_NAME rm /root/$MARIADB_TAR
# docker exec $CONTAINER_NAME rm /root/$MOVIE_SVC_TAR

echo "Images restored in the mydind container registry successfully."
