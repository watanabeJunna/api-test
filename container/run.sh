#!/usr/bin/bash

CONTAINER_NAME="$(openssl rand -hex 16)"

docker run -d --rm --name ${CONTAINER_NAME} -p 27017:27017 mongo:3.4.22-xenial

ID="$(docker container ls -a | grep ${CONTAINER_NAME} | awk '{print $1}')"

[ -z "${ID}" ] && {
    echo "[ERR]: Container is not running" 2>&1
    exit 1
}

IP="$(docker exec -it ${ID} hostname -i)"

docker run --rm -it -w //tmp -v /`pwd`://tmp node:10.16.2-alpine npm i && ./node_modules/ts-node/dist/bin.js index.ts ${IP}

docker stop ${CONTAINER_NAME}