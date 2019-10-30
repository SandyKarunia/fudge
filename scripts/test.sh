#!/usr/bin/env bash

set -eo pipefail

print_title() {
  echo -e "\033[0;36m"
  echo "===================="
  echo "$1"
  echo "===================="
  echo -e "\033[0m"
}

print_title "Running tests..."
go test -cover -race -coverprofile .coverage.out ./...
go tool cover -func .coverage.out

print_title "Running go fmt..."
goFmtOutput="$(go fmt ./...)"
if [ -n "${goFmtOutput}" ]
then
  echo "go fmt fails on the following files:"
  echo "${goFmtOutput}"
  exit 1
fi

print_title "Running golint..."
# find all folders under current directory, excluding hidden items and /mocks folders
foldersToLint=($(find . -type d -not -path '*/\.*' | grep -v /mocks))
for t in "${foldersToLint[@]}"
do
  echo "$t"
done
"${GOPATH}"/bin/golint -set_exit_status "${foldersToLint[@]}"

print_title "Uploading coverage result to coveralls..."
"${GOPATH}"/bin/goveralls -coverprofile=.coverage.out -service=circle-ci -repotoken="${COVERALLS_TOKEN}"