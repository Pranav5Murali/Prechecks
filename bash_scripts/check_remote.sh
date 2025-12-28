#!/bin/bash

# Get the password from the argument
SSH_PASSWORD="$1"

# Define the remote user and IP address
REMOTE_USER="user1"
REMOTE_HOST="ubuntu-target" 

# Run the commands on the remote machine using sshpass
echo "Connecting to remote machine..."
sshpass -p "$SSH_PASSWORD" ssh -o StrictHostKeyChecking=no "$REMOTE_USER@$REMOTE_IP" <<EOF
echo "Running commands on remote machine:"
whoami
ip a
EOF
