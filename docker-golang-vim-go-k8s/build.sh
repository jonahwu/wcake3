#!/bin/bash
REPO=172.16.155.136:5000
REPO=${REPO_JENKINS:-$REPO}

CONTAINER="dockergolangvimgo"
CONTAINER=${CONTAINER_JENKINS:-$CONTAINER}

#TAG=$(git rev-parse --short HEAD)
TAG="1.11"


DOCKER_IMAGE=$REPO/$CONTAINER:$TAG

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
echo $DIR
#BUILDROOT=$DIR/..
BUILDROOT=./

# Build dockerbas
#cmdsource="$DIR/appscript.sh"
#echo $cmdsource
#eval $cmdsource
#cmd="docker build -t $CONTAINER:$TAG -f $DIR/Dockerfile $BUILDROOT"
# $JOB_NAME is provided by Jenkins
cmd="docker build -t $DOCKER_IMAGE -f $DIR/Dockerfile $BUILDROOT"
echo $cmd
eval $cmd
