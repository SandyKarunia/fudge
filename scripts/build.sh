#!/usr/bin/env bash

set -eo pipefail

releaseTag="${FUDGE_BUILD_ID}"
if [ -n "${FUDGE_BUILD_TAG}" ]
then
  releaseTag="${FUDGE_BUILD_TAG}"
fi

echo "Release Tag: ${releaseTag}"

go build -o "${FUDGE_BUILD_DIR}"/fudge-"${releaseTag}" -v

echo "ls build dir:"
ls "${FUDGE_BUILD_DIR}"
