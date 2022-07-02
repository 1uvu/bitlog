#!/bin/bash

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
  if [ -d "${pwd}/${CONTAINER_NAME}" ]; then
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
    echo "clean temp files..."
    sudo rm -rf "${pwd}"/"${temp_dir}"/
  fi
  exit 0
}

function main() {
  echo "copy .env tmpl"
  cp "${pwd}/tmpl/.env.tmpl" "${temp_dir}/.env"
  echo "copy btcd.conf tmpl"
  mkdir "${pwd}/${temp_dir}/${BTCD_ROOT_DIR}"
  cp "${pwd}/tmpl/btcd.conf.tmpl" "${pwd}/${temp_dir}/${BTCD_ROOT_DIR}/btcd.conf"

  echo "copy tmpl to peer ${CONTAINER_NAME}"
  mkdir -p "${pwd}/peers/${CONTAINER_NAME}/"
  cp "${pwd}/${temp_dir}/.env" "${pwd}/peers/${CONTAINER_NAME}/"
  cp -r "${pwd}/${temp_dir}/${BTCD_ROOT_DIR}" "${pwd}/peers/${CONTAINER_NAME}/"
  cp "${pwd}/docker-compose.yaml.tmpl" "${pwd}/peers/${CONTAINER_NAME}/docker-compose.yaml"

  echo "source env"
  cd "${pwd}/peers/${CONTAINER_NAME}" || exit
  source .env

  echo "bootstrap peer"
  docker-compose up -d

  rm -rf "${pwd}/${temp_dir}"
  echo "bootstrap success"
  exitWithError
}

precheck

prepare

main
