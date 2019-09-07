#!/usr/bin/env bash

echo "Running go fmt..."
goFmtOutput="$(go fmt ./...)"
if [ -n "${goFmtOutput}" ]
then
  echo "go fmt fails on the following files:"
  echo ${goFmtOutput}
  exit 1
fi

echo "Running golint..."
goLintOutput=$("${GOPATH}"/bin/golint -set_exit_status ./...)