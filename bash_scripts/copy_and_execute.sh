#!/bin/bash

ssh_pwd="$1"
TARGET_HOST="192.168.1.105"
TARGET_USER="user1"

echo "Copying python script to target container..."
sshpass -p "$ssh_pwd" \
scp -o StrictHostKeyChecking=no \
python_scripts/info1.py \
${TARGET_USER}@${TARGET_HOST}:/home/user1/target/

echo "Executing python script on target container..."
sshpass -p "$ssh_pwd" \
ssh -o StrictHostKeyChecking=no \
${TARGET_USER}@${TARGET_HOST} \
'cd /home/user1/target && python3 info1.py'
