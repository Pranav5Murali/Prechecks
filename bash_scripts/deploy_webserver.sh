#!/bin/bash

# Get the password from the argument
SSH_PASSWORD="$1"

# Define the remote user and IP address
REMOTE_USER="user1"
REMOTE_IP="192.168.1.105"

# Path to the Go source code
LOCAL_SOURCE="./go_scripts/deploy_webserver.go"
REMOTE_SOURCE="/home/$REMOTE_USER/deploy_webserver.go"

# Check if LOCAL_SOURCE exists
if [ ! -f "$LOCAL_SOURCE" ]; then
  echo "Error: Local source file $LOCAL_SOURCE does not exist."
  exit 1
fi

echo "Transferring Go source code to the remote machine..."
sshpass -p "$SSH_PASSWORD" scp "$LOCAL_SOURCE" "$REMOTE_USER@$REMOTE_IP:$REMOTE_SOURCE"
if [ $? -ne 0 ]; then
  echo "Error: Failed to transfer Go source code to the remote machine."
  exit 1
fi

echo "Compiling and running the Go program on the remote machine..."
sshpass -p "$SSH_PASSWORD" ssh -tt -o StrictHostKeyChecking=no "$REMOTE_USER@$REMOTE_IP" <<EOF
  # Ensure Go is installed
  if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed on the remote machine."
    exit 1
  fi

  # Compile the Go program
  go build -o /tmp/deploy_webserver $REMOTE_SOURCE
  if [ $? -ne 0 ]; then
    echo "Error: Failed to compile the Go program."
    exit 1
  fi

  # Make the binary executable and run it
  chmod +x /tmp/deploy_webserver
  /tmp/deploy_webserver
  if [ $? -ne 0 ]; then
    echo "Error: Failed to execute the Go program."
    exit 1
  fi
EOF

if [ $? -ne 0 ]; then
  echo "Error: An issue occurred on the remote machine."
  exit 1
fi

echo "Go-based deployment completed successfully!"
