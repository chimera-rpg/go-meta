# Chimera golang packages

All of these packages presume the usage of [gb](https://getgb.io/). The following assumptions are made:

  * Source for each package is stored in `src/`
  * Produced binaries are stored in `bin/`
  * Run-time data is stored in `share/chimera`
  * Programs are expected to be run alongside the `share` directory such as by issuing `bin/client` or `bin/server`

## Initial Setup
Chimera-go separates client, server, and common code, along with client and server data files, into different directories that correspond to git repositories. These are:

  * /server
    * https://github.com/chimera-rpg/go-server
    * Server program
  * /client
    * https://github.com/chimera-rpg/go-client
    * Client program
  * /common
    * https://github.com/chimera-rpg/go-common
    * Shared code between the server and the client
  * /share/chimera/archetypes
    * https://github.com/chimera-rpg/archetypes
    * Tiles and their graphics
  * /share/chimera/maps
    * https://github.com/chimera-rpg/maps
    * Maps
  * /share/chimera/client
    * https://github.com/chimera-rpg/client-data
    * Data assets for the client, such as UI graphics, sounds, and otherwise.

Optionally, for graphics development, you will likely wish to check out `https://github.com/chimera-rpg/assets` to acquire template graphics, examples, and similar.

### Get all

    ./clone.sh

### Server setup

    git clone ...chimera-go/server src/server
    git clone ...chimera-go/common src/common
    git clone ...chimera-go/archetypes share/chimera/archetypes
    git clone ...chimera-go/maps share/chimera/maps

### Client setup

    git clone ...chimera-go/client src/client
    git clone ...chimera-go/common src/common
    git clone ...chimera-go/client-data share/chimera/client

## Building

Building can be done by issuing `gb build` to build all software or `gb build package_name` to build a particular package located in the `src/` subdirectory.

    gb build
    ... or ...
    gb build client
    gb build server
