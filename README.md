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

### Gobl-based update/build/watch

To see a list of all available tasks:

```
$ go run .
```

For example, to acquire all development repositories:

```
$ go run . updateAll
```

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

### Gobl Build

See the various build tasks via:

```
$ go run .
```

To build the client:

```
$ go run . buildClient
```

To watch the server sources and recompile on change:

```
$ go run . watchServer
```

### Manual
Building can be done by entering into a given module and issuing `go build -o ../../bin/package_name`.

## Additional Setup
### Client<->Server TLS Encryption
Chimera supports using an encrypted TCP connection. This can be done by generating the necessary `server.crt` and `server.key` files through the following commands:

```
openssl genrsa -out server.key 2048
openssl ecparam -genkey -name secp384r1 -out server.key
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```

At the moment these files are expected to exist in `etc/chimera/` directory.
