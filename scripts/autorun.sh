#!/bin/bash
echo "KILL API"
sudo pkill ecoapigorm
sudo systemctl stop ecoapigorm

echo "ADD EXECUTION PERMISSIONS"
chmod +x /home/ubuntu/go/src/github.com/nikola43/ecoapigorm/ecoapigorm

echo "RUN API"
source /home/ubuntu/go/src/github.com/nikola43/ecoapigorm/.env
sudo systemctl start ecoapigorm
exit
