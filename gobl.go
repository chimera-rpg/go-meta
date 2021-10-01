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
		Parallel("updateMeta", "updateServer", "updateCommon", "updateEditor", "updateEditorAssets", "updateArchetypes", "updateMaps", "updateClient", "updateClientAssets")

	repos := map[string][2]string{
		"updateMeta":         {"./", "https://github.com/chimera-rpg/go-meta"},
		"updateCommon":       {"src/go-common", "https://github.com/chimera-rpg/go-common"},
		"updateServer":       {"src/go-server", "https://github.com/chimera-rpg/go-server"},
		"updateEditor":       {"src/go-editor", "https://github.com/chimera-rpg/go-editor"},
		"updateEditorAssets": {"share/chimera/editor", "https://github.com/chimera-rpg/editor-data"},
		"updateArchetypes":   {"share/chimera/archetypes", "https://github.com/chimera-rpg/archetypes"},
		"updateMaps":         {"share/chimera/maps", "https://github.com/chimera-rpg/maps"},
		"updateClient":       {"src/go-client", "https://github.com/chimera-rpg/go-client"},
		"updateClientAssets": {"share/chimera/client", "https://github.com/chimera-rpg/client-data"},
	}
	for taskName, repo := range repos {
		func(taskName string, repo [2]string) {
			Task(taskName).
				Exists(repo[0]).
				Catch(func(err error) error {
					cmd := exec.Command("git", "clone", repo[1], repo[0])
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
		"buildEditor": {"src/go-editor", "../../bin/editor" + exe},
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
		Watch("src/go-client/*.go", "src/go-client/*/*.go").
		Signaler(SigQuit).
		Run("buildClient").
		Run("runClient")

	Task("watchEditor").
		Watch("src/go-editor/*.go", "src/go-editor/*/*.go", "src/go-editor/*/*/*.go").
		Signaler(SigQuit).
		Run("buildEditor").
		Run("runEditor")

	Task("runServer").
		Exec("./bin/server"+exe, "--no-prompt")

	Task("runEditor").
		Exec("./bin/editor" + exe)

	Task("runClient").
		Exec("./bin/client" + exe)

	Go()
}
