#!/bin/bash
cd /home/ubuntu/go/src/github.com/nikola43/ecoapigorm

echo "KILL API"
sudo pkill ecoapigorm
sudo systemctl stop ecoapigorm

echo "RUN API"
source /home/ubuntu/go/src/github.com/nikola43/ecoapigorm/.env
sudo systemctl start ecoapigorm
exit
