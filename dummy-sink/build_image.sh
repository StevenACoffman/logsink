#!/bin/bash
set -x

IMAGE_NAME=dummy-logsink
REPOSITORY_NAMESPACE=${1:-stevenacoffman}

REPOSITORY="${REPOSITORY_NAMESPACE}/${IMAGE_NAME}"

docker build \
    -t "${REPOSITORY}:latest" \
    -f Dockerfile .

docker push "${REPOSITORY}"
