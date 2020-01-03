#!/bin/bash
:<<BATCH
@echo off

if /i "%1"=="client" goto build_client
if /i "%1"=="server" goto build_server
goto prompt

:build_client
echo "Building client..."
cd src\go-client
go build -o ..\..\bin\client
goto:eof

:build_server
  echo "Building server..."
cd src\go-server
go build -o ..\..\bin\server
goto:eof

:prompt
echo Issue "client" or "server" to build either.
goto:eof
BATCH

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
