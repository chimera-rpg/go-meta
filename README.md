# Chimera golang packages

All of the packages pulled from this script are either go modules or are data repositories. The following assumptions are made:

  * Source for each package is stored in `src/`
  * Produced binaries are stored in `bin/`
  * Run-time data is stored in `share/chimera`
  * Programs are expected to be run alongside the `share` directory such as by issuing `bin/client` or `bin/server`

## Initial Setup
Chimera-go separates client, server, and common code, along with client and server data files, into different directories that correspond to git repositories. These are:

  * /src/go-server
    * https://github.com/chimera-rpg/go-server
    * Server program
  * /src/go-client
    * https://github.com/chimera-rpg/go-client
    * Client program
  * /src/go-common
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

Building can be done by entering into a given module and issuing `go build -o ../../bin/package_name`.
