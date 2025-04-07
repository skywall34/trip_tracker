#!/bin/bash

# Exit if any command fails
set -e

# === CONFIG ===
IMAGE_NAME="trip_tracker"
TAG="latest"
REGISTRY="docker.io"  # Change to ghcr.io or your registry
USERNAME="skywall34"  # Change to your Docker Hub or registry username
FULL_IMAGE_NAME="$REGISTRY/$USERNAME/$IMAGE_NAME:$TAG"

# === BUILD ===
echo "Building Docker image..."
docker build -t "$IMAGE_NAME:$TAG" .

# === TAG ===
echo "Tagging image as $FULL_IMAGE_NAME"
docker tag "$IMAGE_NAME:$TAG" "$FULL_IMAGE_NAME"

# === LOGIN ===
echo "Logging in to $REGISTRY"
docker login "$REGISTRY"

# === PUSH ===
echo "Pushing image to $FULL_IMAGE_NAME"
docker push "$FULL_IMAGE_NAME"

echo "✅ Done!"