#!/bin/bash
:<<BATCH
@echo off

IF EXIST src\go-client (
  echo "Updating client..."
  cd src\go-client
  git pull
  cd ..\..
) ELSE (
  git clone https://github.com/chimera-rpg/go-client src\go-client
)
IF EXIST src\go-server (
  echo "Updating server..."
  cd src\go-server
  git pull
  cd ..\..
) ELSE (
  git clone https://github.com/chimera-rpg/go-server src\go-server
)
IF EXIST src\go-editor (
  echo "Updating editor..."
  cd src\go-editor
  git pull
  cd ..\..
) ELSE (
  git clone https://github.com/chimera-rpg/go-editor src\go-editor
)
IF EXIST src\go-common (
  echo "Updating common..."
  cd src\go-common
  git pull
  cd ..\..
) ELSE (
  git clone https://github.com/chimera-rpg/go-common src\go-common
)
IF EXIST share/chimera/client (
  echo "Updating client assets..."
  cd share\chimera\client
  git pull
  cd ..\..\..
) ELSE (
  git clone https://github.com/chimera-rpg/client-data share\chimera\client
)
IF EXIST share/chimera/archetypes (
  echo "Updating archetypes..."
  cd share\chimera\archetypes
  git pull
  cd ..\..\..
) ELSE (
  git clone https://github.com/chimera-rpg/archetypes share\chimera\archetypes
)
IF EXIST share/chimera/maps (
  echo "Updating maps..."
  cd share\chimera\maps
  git pull
  cd ..\..\..
) ELSE (
  git clone https://github.com/chimera-rpg/maps share\chimera\maps
)
IF EXIST share/chimera/editor (
  echo "Updating editor assets..."
  cd share\chimera\editor
  git pull
  cd ..\..\..
) ELSE (
  git clone https://github.com/chimera-rpg/editor-data share\chimera\editor
)


goto:eof
BATCH

if [ -d "src/go-client" ]
then
  echo "Updating client..."
  cd src/go-client
  git pull
  cd ../..
else
  echo "Cloning client..."
  git clone https://github.com/chimera-rpg/go-client src/go-client
fi

if [ -d "src/go-server" ]
then
  echo "Updating server..."
  cd src/go-server
  git pull
  cd ../..
else
  echo "Cloning server..."
  git clone https://github.com/chimera-rpg/go-server src/go-server
fi

if [ -d "src/go-editor" ]
then
  echo "Updating editor..."
  cd src/go-editor
  git pull
  cd ../..
else
  echo "Cloning editor..."
  git clone https://github.com/chimera-rpg/go-editor src/go-editor
fi

if [ -d "src/go-common" ]
then
  echo "Updating common..."
  cd src/go-common
  git pull
  cd ../..
else
  echo "Cloning common..."
  git clone https://github.com/chimera-rpg/go-common src/go-common
fi

if [ -d "share/chimera/client" ]
then
  echo "Updating client assets..."
  cd share/chimera/client
  git pull
  cd ../../..
else
  echo "Cloning client assets..."
  git clone https://github.com/chimera-rpg/client-data share/chimera/client
fi

if [ -d "share/chimera/archetypes" ]
then
  echo "Updating server archetypes..."
  cd share/chimera/archetypes
  git pull
  cd ../../..
else
  echo "Cloning server archetypes..."
  git clone https://github.com/chimera-rpg/archetypes share/chimera/archetypes
fi

if [ -d "share/chimera/maps" ]
then
  echo "Updating server maps..."
  cd share/chimera/maps
  git pull
  cd ../../..
else
  echo "Cloning server maps..."
  git clone https://github.com/chimera-rpg/maps share/chimera/maps
fi

if [ -d "share/chimera/editor" ]
then
  echo "Updating editor assets..."
  cd share/chimera/editor
  git pull
  cd ../../..
else
  echo "Cloning editor assets..."
  git clone https://github.com/chimera-rpg/editor-data share/chimera/editor
fi

