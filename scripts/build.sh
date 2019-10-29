#!/usr/bin/env bash

set -eo pipefail

releaseTag="${CI_PIPELINE_ID}"
if [ -n "${CI_COMMIT_TAG}" ]
then
  releaseTag="${CI_COMMIT_TAG}"
fi

echo "Release Tag: ${releaseTag}"

go build -o "${BUILD_DIST_FOLDER}"/fudge-"${releaseTag}" -v
