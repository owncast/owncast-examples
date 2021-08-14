#!/bin/sh

docker build . -t owncast-demobot
docker tag owncast-demobot:latest ghcr.io/owncast/owncast-demobot:latest
docker push ghcr.io/owncast/owncast-demobot:latest

