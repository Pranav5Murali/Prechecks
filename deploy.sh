#!/bin/bash

# Configuration
SLAVE_USER="user"
SLAVE_HOST="192.168.1.105"
SLAVE_PWD=$1
APP_PATH="/home/user/Python_github"

# Step 1: Copy project to the slave
echo "Copying project to slave machine..."
sshpass -p "$SLAVE_PWD" scp -r . $SLAVE_USER@$SLAVE_HOST:$APP_PATH

# Step 2: SSH into slave and deploy the app
echo "Deploying app on slave machine..."
sshpass -p "$SLAVE_PWD" ssh $SLAVE_USER@$SLAVE_HOST bash <<EOF
    cd $APP_PATH
    docker-compose down || true
    docker-compose up --build -d
EOF

echo "Deployment complete! Flask app should be accessible at http://$SLAVE_HOST:5001"
