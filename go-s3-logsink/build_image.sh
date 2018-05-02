#!/bin/bash
set -x

IMAGE_NAME=go-s3-logsink
REPOSITORY_NAMESPACE=${1:-stevenacoffman}

REPOSITORY="${REPOSITORY_NAMESPACE}/${IMAGE_NAME}"

docker build \
    -t "${REPOSITORY}:latest" \
    -f Dockerfile .

docker push "${REPOSITORY}"
