#!/bin/bash
cd /home/ubuntu/go/src/github.com/nikola43/ecoapi

echo "KILL API"
sudo pkill ecoapi
sudo systemctl stop ecoapi

sudo pkill ecoapi
echo "RUN API"
#./ecoapi &
sudo systemctl start ecoapi
exit
