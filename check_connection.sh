#!/bin/bash
ADDR=http://localhost:8080
ADDR=$NODE_IP:$NODE_PORT
while true; do
    curl $ADDR;
    sleep 1;
done;
