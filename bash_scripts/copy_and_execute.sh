#!/bin/bash
ssh_pwd="$1"

echo "Connecting to the remote machine and copy script"
sshpass -p "$ssh_pwd" \
scp python_scripts/info1.py user1@192.168.1.105:/home/user1/target

echo "executing the python script on the remote machine"
sshpass -p "$ssh_pwd" \
ssh user1@192.168.1.105 \
'cd /home/user1/target && python3 info1.py'
