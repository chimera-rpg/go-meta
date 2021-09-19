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
	}

	// Create our build tasks.
	builds := map[string][2]string{
		"buildServer": {"src/go-server", "../../bin/server" + exe},
		"buildEditor": {"src/go-editor", "../../bin/editor" + exe},
		"buildClient": {"src/go-client", "../../bin/client" + exe},
	}
	for taskName, build := range builds {
		Task(taskName).
			Chdir(build[0]).
			Exec("go", "build", "-v", "-o", build[1])
	}

	Task("updateAll").
		Run("updateMeta").
		Run("updateServer").
		Run("updateCommon").
		Run("updateEditor").
		Run("updateEditorAssets").
		Run("updateArchetypes").
		Run("updateMaps").
		Run("updateClient").
		Run("updateClientAssets")

	Task("buildAll").
		Run("buildServer").
		Run("buildEditor").
		Run("buildClient")

	Task("watchServer").
		Watch("src/go-server/*.go", "src/go-server/*/*.go", "src/go-server/*/*/*.go").
		Run("buildServer").
		Run("runServer")

	Task("watchClient").
		Watch("src/go-client/*.go", "src/go-client/*/*.go").
		Run("buildClient").
		Run("runClient")

	Task("watchEditor").
		Watch("src/go-editor/*.go", "src/go-editor/*/*.go", "src/go-editor/*/*/*.go").
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
