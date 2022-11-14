package main

import (
	"os/exec"
	"runtime"

	. "github.com/kettek/gobl"
)

func main() {
	var exe string
	if runtime.GOOS == "windows" {
		exe = ".exe"
	}

	// Create our repo tasks.
	Task("updateAll").
		Parallel("updateMeta", "updateServer", "updateEditor", "updateArchetypes", "updateMaps", "updateAudio", "updateClient", "updateClientAssets")

	repos := map[string][2]string{
		"updateMeta":         {"./", "github.com/chimera-rpg/go-meta"},
		"updateServer":       {"src/go-server", "github.com/chimera-rpg/go-server"},
		"updateEditor":       {"src/chedit", "github.com/chimera-rpg/chedit"},
		"updateArchetypes":   {"share/chimera/archetypes", "github.com/chimera-rpg/archetypes"},
		"updateAudio":        {"share/chimera/audio", "github.com/chimera-rpg/audio"},
		"updateMaps":         {"share/chimera/maps", "github.com/chimera-rpg/maps"},
		"updateClient":       {"src/go-client", "github.com/chimera-rpg/go-client"},
		"updateClientAssets": {"share/chimera/client", "github.com/chimera-rpg/client-data"},
	}
	for taskName, repo := range repos {
		func(taskName string, repo [2]string) {
			Task(taskName).
				Exists(repo[0]).
				Catch(func(err error) error {
					cmd := exec.Command("git", "clone", "https://"+repo[1], repo[0])
					err = cmd.Run()
					return err
				}).
				Chdir(repo[0]).
				Exec("git", "pull").
				Result(func(r interface{}) {})
		}(taskName, repo)
	}

	// Create our build tasks.
	Task("buildAll").
		Run("buildServer").
		Run("buildEditor").
		Run("buildClient")

	builds := map[string][2]string{
		"buildServer": {"src/go-server", "../../bin/server" + exe},
		"buildClient": {"src/go-client", "../../bin/client" + exe},
	}
	for taskName, build := range builds {
		func(taskName string, build [2]string) {
			Task(taskName).
				Chdir(build[0]).
				Exec("go", "build", "-v", "-o", build[1])
		}(taskName, build)
	}

	Task("watchServer").
		Watch("src/go-server/*.go", "src/go-server/*/*.go", "src/go-server/*/*/*.go").
		Signaler(SigQuit).
		Run("buildServer").
		Run("runServer")

	Task("watchClient").
		Watch("src/go-client/*.go", "src/go-client/*/*.go", "src/go-client/*/*/*.go", "src/go-client/*/*/*/*.go").
		Signaler(SigQuit).
		Run("buildClient").
		Run("runClient")

	Task("runServer").
		Exec("./bin/server"+exe, "--no-prompt")

	Task("runEditor").
		Exec("./bin/editor" + exe)

	Task("runClient").
		Exec("./bin/client" + exe)

	serverModule := ""
	Task("getServerSHA").
		Chdir("src/go-server").
		Exec("git", "rev-parse", "HEAD").
		Result(func(i interface{}) {
			serverSHA := i.(string)
			serverSHA = serverSHA[:len(serverSHA)-1]
			serverModule = repos["updateServer"][1] + "@" + serverSHA
		})

	Task("updateDeps").
		Run("getServerSHA").
		Chdir("src/go-client").
		Exec("go", "get", "-v", "-u", &serverModule).
		Chdir("../chedit").
		Exec("go", "get", "-v", "-u", &serverModule)

	Go()
}
