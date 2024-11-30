#!/bin/bash

# Get the password from the argument
SSH_PASSWORD="$1"

# Define the remote user and IP address
REMOTE_USER="user1"
REMOTE_IP="192.168.1.105"

# Path to the Go source code
LOCAL_SOURCE="./go_scripts/deploy_webserver.go"
REMOTE_SOURCE="/tmp/deploy_webserver.go"

echo "Transferring Go source code to the remote machine..."
sshpass -p "$SSH_PASSWORD" scp "$LOCAL_SOURCE" "$REMOTE_USER@$REMOTE_IP:$REMOTE_SOURCE"

echo "Compiling and running the Go program on the remote machine..."
sshpass -p "$SSH_PASSWORD" ssh -o StrictHostKeyChecking=no "$REMOTE_USER@$REMOTE_IP" <<EOF
  go build -o /tmp/deploy_webserver $REMOTE_SOURCE
  chmod +x /tmp/deploy_webserver
  /tmp/deploy_webserver
EOF

echo "Go-based deployment completed successfully!"
