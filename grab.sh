#!/bin/bash

if [ -d "src/go-client" ]
then
  echo "Updating client..."
  cd src/go-client && git pull && cd ../..
else
  echo "Cloning client..."
  git clone https://github.com/chimera-rpg/go-client src/go-client
fi

if [ -d "src/go-server" ]
then
  echo "Updating server..."
  cd src/go-server && git pull && cd ../..
else
  echo "Cloning server..."
  git clone https://github.com/chimera-rpg/go-server src/go-server
fi

if [ -d "src/go-common" ]
then
  echo "Updating common..."
  cd src/go-common && git pull && cd ../..
else
  echo "Cloning common..."
  git clone https://github.com/chimera-rpg/go-common src/go-common
fi

if [ -d "share/chimera/client" ]
then
  echo "Updating client assets..."
  cd share/chimera/client && git pull && cd ../../..
else
  echo "Cloning client assets..."
  git clone https://github.com/chimera-rpg/client-data share/chimera/client
fi

if [ -d "share/chimera/archetypes" ]
then
  echo "Updating server archetypes..."
  cd share/chimera/archetypes && git pull && cd ../../..
else
  echo "Cloning server archetypes..."
  git clone https://github.com/chimera-rpg/archetypes share/chimera/archetypes
fi

if [ -d "share/chimera/maps" ]
then
  echo "Updating server maps..."
  cd share/chimera/maps && git pull && cd ../../..
else
  echo "Cloning server maps..."
  git clone https://github.com/chimera-rpg/maps share/chimera/maps
fi
