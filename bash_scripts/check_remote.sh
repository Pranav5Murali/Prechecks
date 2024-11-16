#!/bin/bash

# Get the password from the argument
SSH_PASSWORD="$1"

# Define the remote user and IP address
REMOTE_USER="karthik"
REMOTE_IP="192.168.1.105"

# Run the commands on the remote machine using sshpass
echo "Connecting to remote machine..."
sshpass -p "$SSH_PASSWORD" ssh "$REMOTE_USER@$REMOTE_IP" <<EOF
echo "Running commands on remote machine:"
whoami
ip a
uname -a
EOF
