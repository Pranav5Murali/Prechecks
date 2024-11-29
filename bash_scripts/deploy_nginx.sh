#!/bin/bash

# Get the password from the argument
SSH_PASSWORD="$1"

# Define the remote user and IP address
REMOTE_USER="user1"
REMOTE_IP="192.168.1.105"

# Deploy the Nginx container on the slave machine
echo "Connecting to the remote machine and deploying Nginx..."
sshpass -p "$SSH_PASSWORD" ssh -o StrictHostKeyChecking=no "$REMOTE_USER@$REMOTE_IP" <<EOF
  echo "Pulling the latest Nginx image..."
  docker pull nginx:latest

  echo "Stopping and removing any existing Nginx container..."
  docker stop sample-nginx || true
  docker rm sample-nginx || true

  echo "Running a new Nginx container..."
  docker run -d -p 8080:80 --name sample-nginx nginx:latest

  echo "Nginx container deployed successfully!"
EOF
