#!/usr/bin/env bash

set -eo pipefail

print_title() {
  echo -e "\033[0;36m"
  echo "===================="
  echo "$1"
  echo "===================="
  echo -e "\033[0m"
}

# We are adding temp test files because go test will ignore the coverage for packages with no tests
# to get the true coverage, we need to put test files in every packages
# find all directories that DON'T have test files, excludes hidden and mocks directories
dirsNoTest=$(
  find . \
  -not -path '*/\.*' \
  -not -path '*/mocks' \
  -not -path '*/scripts' \
  -not -path '\.' \
  -type d '!' \
  -exec sh -c 'ls -1 "{}"|egrep -i -q ".*_test\.go$"' ';' \
  -print
)
tempTestFileName="temporary_only_do_not_commit_test.go"

print_title "Creating temporary test files in the directories with no test file..."
# shellcheck disable=SC2068
# disable the double-quotes check since we want it to split by newline
for dir in ${dirsNoTest[@]}
do
  # for each directory, create a temporary test file so we can calculate true coverage later
  tempTestPath="${dir}/${tempTestFileName}"
  tempTestContent="package ${dir##*/}"
  touch "${tempTestPath}"
  echo "${tempTestContent}" > "${tempTestPath}"
  echo "Created ${tempTestPath} with content '${tempTestContent}'"
done

print_title "Running tests..."
go test -cover -race -coverprofile .coverage.out ./...
go tool cover -func .coverage.out

print_title "Removing temporary test files on ${#dirsNoTest[@]} directories..."
# shellcheck disable=SC2068
# disable the double-quotes check since we want it to split by newline
for dir in ${dirsNoTest[@]}
do
  # for each directory, remove the created temporary test file
  tempTestPath="${dir}/${tempTestFileName}"
  rm "${tempTestPath}"
  echo "Removed ${tempTestPath}"
done

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

if [ "${CIRCLECI}" == "true" ]
then
  print_title "Uploading coverage result to coveralls..."
  "${GOPATH}"/bin/goveralls -coverprofile=.coverage.out -service=circle-ci -repotoken="${COVERALLS_TOKEN}"
fi