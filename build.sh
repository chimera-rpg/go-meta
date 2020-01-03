#!/bin/bash

build_client()
{
  echo "Building client..."
  cd src/go-client
  go build -o ../../bin/client
  cd ../..
}

build_server() {
  echo "Building server..."
  cd src/go-server
  go build -o ../../bin/server
  cd ../..
}

if [ "$1" = "client" ]
then
  build_client
elif [ "$1" = "server" ]
then
  build_server
elif [ "$1" = "" ]
then
  build_client
  build_server
fi
