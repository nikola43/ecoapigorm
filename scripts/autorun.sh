#!/bin/bash
cd /home/ubuntu/go/src/github.com/nikola43/ecoapigorm

echo "KILL API"
sudo pkill ecoapigorm
sudo systemctl stop ecoapigorm

echo "RUN API"
sudo systemctl start ecoapigorm
exit
