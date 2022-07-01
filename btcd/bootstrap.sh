#!/bin/bash

pwd=$PWD
cd "$PWD"/docker-compose/ || exit 0
echo "copy env example to env"
cp .env.example .env
source .env

if [ -z "${BTCD_ROOT_DIR}" ]; then
  echo "error: peer root dir is empty"
  exit 255
fi
if [ -d "${BTCD_ROOT_DIR}" ]; then
  echo "error: peer root dir has existed"
  exit 254
fi
mkdir -p "${BTCD_ROOT_DIR}"
echo "copy peer example to peer ${CONTAINER_NAME}"
cp peers/peer-example/btcd.conf "${BTCD_ROOT_DIR}"/

echo "bootstrap peer"
docker-compose -f docker-compose-tmpl.yaml up -d

echo "bootstrap success"
cd "$pwd" || exit 0
