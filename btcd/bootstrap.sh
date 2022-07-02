#!/bin/bash

subnet_name=bitlog_peer_network
temp_dir=temp
pwd=${PWD}/docker-compose

function precheck() {
  if [ ! -f "${pwd}/tmpl/.env.tmpl" ]; then
    exitWithError ".env tmpl not prepare"
  fi
  if [ ! -f "${pwd}/tmpl/btcd.conf.tmpl" ]; then
    exitWithError "btcd.conf.tmpl not prepared"
  fi
  source "${pwd}/tmpl/.env.tmpl"
  if [ -z "${CONTAINER_NAME}" ]; then
    exitWithError "peer name not set, please set"
  fi
  if [ -z "${BTCD_ROOT_DIR}" ]; then
    exitWithError "btcd root dir not set, please set"
  fi
  if [ -d "${pwd}/peers/${CONTAINER_NAME}" ]; then
    exitWithError "peer container dir has existed, please rm or rename peer"
  fi
}

function prepare() {
  cd "${pwd}"/ || exit 0
  rm -rf "${temp_dir}"
  mkdir "${temp_dir}"
}

function exitWithError() {
  errorMsg=$1
  if [ -n "${errorMsg}" ]; then
    echo "error: ${errorMsg}"
  fi
  if [ -d "${temp_dir}" ]; then
    echo "clean temp files"
    sudo rm -rf "${pwd}"/"${temp_dir}"/
  fi
  exit 0
}

function bootstrap() {
  docker network create --driver=bridge --subnet="${SUBNET_CIDR}" "${subnet_name}"
  docker-compose up -d
}

function main() {
  echo "precheck process"
  precheck

  echo "prepare process"
  prepare

  echo "copy .env tmpl"
  cp "${pwd}/tmpl/.env.tmpl" "${temp_dir}/.env"
  echo "copy btcd.conf tmpl"
  cp "${pwd}/tmpl/btcd.conf.tmpl" "${pwd}/${temp_dir}/btcd.conf"

  echo "copy tmpl to peer ${CONTAINER_NAME}"
  mkdir -p "${pwd}/peers/${CONTAINER_NAME}/"
  cp "${pwd}/${temp_dir}/.env" "${pwd}/peers/${CONTAINER_NAME}/"
  mkdir "${pwd}/peers/${CONTAINER_NAME}/${BTCD_ROOT_DIR}/"
  cp "${pwd}/${temp_dir}/btcd.conf" "${pwd}/peers/${CONTAINER_NAME}/${BTCD_ROOT_DIR}/"
  cp "${pwd}/docker-compose.yaml.tmpl" "${pwd}/peers/${CONTAINER_NAME}/docker-compose.yaml"

  echo "source env"
  cd "${pwd}/peers/${CONTAINER_NAME}" || exit
  source .env

  echo "bootstrap peer"
  bootstrap

  rm -rf "${pwd}/${temp_dir}"
  echo "bootstrap success"
  exitWithError
}

main
