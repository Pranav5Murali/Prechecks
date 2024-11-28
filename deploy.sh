#!/bin/bash

# Input Variables
ssh_pwd="$1"
SLAVE_USER="user1"
SLAVE_HOST="192.168.1.105"
APP_PATH="/home/user1/Prechecks"
GIT_REPO="https://github.com/Pranav5Murali/Prechecks.git"

echo "Step 1: Clone or update the repository on the slave machine"
sshpass -p "$ssh_pwd" ssh $SLAVE_USER@$SLAVE_HOST bash <<EOF
  if [ -d "$APP_PATH" ]; then
    echo "Repository already exists. Pulling latest changes..."
    cd $APP_PATH && git pull
  else
    echo "Cloning repository..."
    git clone $GIT_REPO $APP_PATH
  fi
EOF

echo "Step 2: Build and run the Docker app on the slave machine"
sshpass -p "$ssh_pwd" ssh $SLAVE_USER@$SLAVE_HOST bash <<EOF
  cd $APP_PATH
  docker-compose down || true
  docker-compose up --build -d
EOF

echo "Deployment complete! Flask app should be accessible at http://$SLAVE_HOST:5001"
