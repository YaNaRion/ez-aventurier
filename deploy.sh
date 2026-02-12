#!/bin/bash

set -e

ENV=$1
VERSION=$2

if [ -z "$ENV" ]; then
    echo "Usage: ./deploy.sh <dev|staging|prod> [version]"
    exit 1
fi

# Checkout specific version if provided
if [ ! -z "$VERSION" ]; then
    git checkout $VERSION
fi

# Build frontend
echo "Building frontend..."
cd client_rust
dx bundle --platform web --release

cd ..

# Deploy
echo "Deploying to $ENV..."
make $ENV

echo "✅ Deployment complete!"
