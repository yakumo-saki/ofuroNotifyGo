#!/bin/bash -eu

SCRIPT_DIR=$(cd $(dirname $0); pwd)

echo "Script dir is ${SCRIPT_DIR}"

read -p "Deploy to AWS Lambda? Yes: Build + Deploy, No: Build only (y/N): " yn

cd ${SCRIPT_DIR}

mkdir -p ${SCRIPT_DIR}/build
rm -rf ${SCRIPT_DIR}/build/*
CGO_ENABLED=0 go build -o ./build/ofuroNotify ofuroNotify.go

cp ${SCRIPT_DIR}/function.json ${SCRIPT_DIR}/build/
#rsync -a -exclude=.git --exclude=package --exclude=deploy.sh --exclude=requirements.txt --exclude=function.zip --exclude=build ${SCRIPT_DIR}/. ${SCRIPT_DIR}/build/

case "$yn" in [yY]*) ;; *) echo "Build complete. (without deploy)" ; exit ;; esac

echo "Deploy to lambda using lambroll"

cd ${SCRIPT_DIR}/build
lambroll deploy --region ap-northeast-1