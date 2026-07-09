#!/bin/bash
set -e

# Get the password from the argument
SSH_PASSWORD="$1"

# Remote target details
REMOTE_USER="user1"
REMOTE_HOST="192.168.1.105"

echo "Connecting to remote machine..."

sshpass -p "$SSH_PASSWORD" ssh \
  -o StrictHostKeyChecking=no \
  "$REMOTE_USER@$REMOTE_HOST" <<EOF
echo "Running commands on remote machine:"
whoami
hostname
ip a
EOF
