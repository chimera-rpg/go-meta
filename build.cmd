#!/bin/bash
:<<BATCH
@echo off

if /i "%1"=="client" goto build_client
if /i "%1"=="server" goto build_server
if /i "%1"=="editor" goto build_editor
goto prompt

:build_client
echo "Building client..."
cd src\go-client
go build -o ..\..\bin\client.exe
goto:eof

:build_server
  echo "Building server..."
cd src\go-server
go build -o ..\..\bin\server.exe
goto:eof

:build_editor
  echo "Building editor..."
cd src\go-editor
go build -o ..\..\bin\editor
goto:eof

:prompt
echo Issue "client", "server", or "editor" to build.
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

build_editor() {
  echo "Building editor..."
  cd src/go-editor
  go build -o ../../bin/editor
  cd ../..
}

if [ "$1" = "client" ]
then
  build_client
elif [ "$1" = "server" ]
then
  build_server
elif [ "$1" = "editor" ]
then
  build_editor
elif [ "$1" = "" ]
then
  build_client
  build_server
  build_editor
fi
