#!/bin/bash
#REPO=docker-reg.emotibot.com.cn:55688
REPO=172.16.155.136:5000
REPO=${REPO_JENKINS:-$REPO}

CONTAINER="k8sinstance"
CONTAINER=${CONTAINER_JENKINS:-$CONTAINER}

TAG=$(git rev-parse --short HEAD)
#TAG=test


DOCKER_IMAGE=$REPO/$CONTAINER:$TAG

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
echo $DIR
#BUILDROOT=$DIR/..
BUILDROOT=$DIR

# Build dockerbas
#cmdsource="$DIR/appscript.sh"
#echo $cmdsource
#eval $cmdsource
#cmd="docker build -t $CONTAINER:$TAG -f $DIR/Dockerfile $BUILDROOT"
# $JOB_NAME is provided by Jenkins
cmd="docker build --no-cache -t $DOCKER_IMAGE -f $DIR/Dockerfile $BUILDROOT"
echo $cmd
eval $cmd
echo "/////////////////////////////"
echo "push result to remote"
cmdpush="docker push $REPO/$CONTAINER:$TAG"
echo $cmdpush
echo "or execute it locally"
cmd1="docker run -it -p 30000:8000 $REPO/$CONTAINER:$TAG"
echo "$cmd1"
#eval $cmd1
