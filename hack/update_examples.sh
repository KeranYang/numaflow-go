#!/bin/bash

function show_help () {
    echo "Usage: $0 [-h|--help | -t|--tag <tag>] (-bpe|--build-push-example <path> | -u|--update <SDK-version | commit-sha | latest >)"
    echo "  -h, --help                   Display help message and exit"
    echo "  -bpe, --build-push-example   Build the Dockerfile of the given example directory path, and push it to the quay.io registry"
    echo "  -t, --tag                    To be optionally used with -bpe. Specify the tag to build with. Default tag: stable"
    echo "  -u, --update                 Update all of the examples to depend on the numaflow-go version with the specified SHA or version"
}

function traverse_examples () {
  find pkg -name "go.mod" | while read -r line;
  do
      dir="$(dirname "${line}")"
      cd "$dir" || exit

      for command in "$@"
      do
        if ! $command; then
          echo "Error: failed $command in $dir" >&2
          exit 1
        fi
      done

      cd ~- || exit
  done
}

if [ $# -eq 0 ]; then
  echo "Error: provide at least one argument" >&2
  show_help
  exit 1
fi

usingHelp=0
usingBuildPushExample=0
usingVersion=0
usingTag=0
version=""
directoryPath=""
tag="stable"

function handle_options () {
  while [ $# -gt 0 ]; do
    case "$1" in
      -h | --help)
        usingHelp=1
        ;;
      -bpe | --build-push-example)
        if [ -z "$2" ]; then
          echo "Directory path not specified." >&2
          show_help
          exit 1
        fi

        usingBuildPushExample=1
        directoryPath=$2
        shift
        ;;
      -t | --tag)
        if [ -z "$2" ]; then
          echo "Tag not specified." >&2
          show_help
          exit 1
        fi

        usingTag=1
        tag=$2
        shift
        ;;
      -u | --update)
        if [ -z "$2" ]; then
          echo "Commit SHA or version not specified." >&2
          show_help
          exit 1
        fi

        usingVersion=1
        version=$2
        shift
        ;;
      *)
        echo "Invalid option: $1" >&2
        show_help
        exit 1
        ;;
    esac
    shift
  done
}

handle_options "$@"

if (( usingBuildPushExample + usingVersion + usingHelp > 1 )); then
  echo "Only one of '-h', '-bpe', or, '-u' is allowed at a time" >&2
  show_help
  exit 1
fi

if (( (usingTag + usingVersion + usingHelp > 1) || (usingTag && usingBuildPushExample == 0) )); then
  echo "Can only use -t with -bpe" >&2
  show_help
  exit 1
fi

if [ -n "$version" ]; then
 echo "Will update to: $version"
fi

if [ -n "$directoryPath" ]; then
 echo "Dockerfile path to use: $directoryPath"
fi

if [ -n "$tag" ] && (( ! usingVersion )) && (( ! usingHelp )); then
 echo "Using tag: $tag"
fi

if (( usingBuildPushExample )); then
   cd "./$directoryPath" || exit
   if ! make image-push TAG="$tag"; then
     echo "Error: failed to run make image-push in $directoryPath" >&2
     exit 1
   fi
elif (( usingVersion )); then
  traverse_examples "go get github.com/numaproj/numaflow-go@$version" "go mod tidy"
elif (( usingHelp )); then
  show_help
fi
