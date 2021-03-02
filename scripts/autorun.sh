#!/bin/bash
cd /home/ubuntu/go/src/github.com/nikola43/ecoapi

echo "UPDATING REPO"
git pull
echo "KILL API"

sudo pkill ecoapi
echo "BUILD API..."

go build
echo "KILL API"

sudo pkill ecoapi
echo "RUN API"
./ecoapi &
