package main

import (
  "os/exec"
	. "github.com/kettek/gobl/gobl"
  "runtime"
  "fmt"
)

func main() {
  var exe string
  if runtime.GOOS == "windows" {
    exe = ".exe"
  }

	Task("updateClient").
    Exists("src/go-client").
    Catch(func(err error) error {
      cmd := exec.Command("git", "clone", "https://github.com/chimera-rpg/go-client", "src/go-client")
      err = cmd.Run()
      return err
    }).
    Chdir("src/go-client").
    Exec("git", "pull").
    Chdir("../../").
    Result(func(r interface{}) {
    })

	Task("updateCommon").
    Exists("src/go-common").
    Catch(func(err error) error {
      cmd := exec.Command("git", "clone", "https://github.com/chimera-rpg/go-common", "src/go-common")
      err = cmd.Run()
      return err
    }).
    Chdir("src/go-common").
    Exec("git", "pull").
    Chdir("../../").
    Result(func(r interface{}) {
    })

	Task("updateServer").
    Exists("src/go-server").
    Catch(func(err error) error {
      cmd := exec.Command("git", "clone", "https://github.com/chimera-rpg/go-server", "src/go-server")
      err = cmd.Run()
      return err
    }).
    Chdir("src/go-server").
    Exec("git", "pull").
    Chdir("../../").
    Result(func(r interface{}) {
    })

	Task("updateEditor").
    Exists("src/go-editor").
    Catch(func(err error) error {
      cmd := exec.Command("git", "clone", "https://github.com/chimera-rpg/go-editor", "src/go-editor")
      err = cmd.Run()
      return err
    }).
    Chdir("src/go-editor").
    Exec("git", "pull").
    Chdir("../../").
    Result(func(r interface{}) {
    })

  Task("updateEditorAssets").
    Exists("share/chimera/editor").
    Catch(func(err error) error {
      cmd := exec.Command("git", "clone", "https://github.com/chimera-rpg/editor-data", "share/chimera/editor")
      err = cmd.Run()
      return err
    }).
    Chdir("share/chimera/editor").
    Exec("git", "pull").
    Chdir("../../../").
    Result(func(r interface{}) {
    })

  Task("updateClientAssets").
    Exists("share/chimera/client").
    Catch(func(err error) error {
      cmd := exec.Command("git", "clone", "https://github.com/chimera-rpg/client-data", "share/chimera/client")
      err = cmd.Run()
      return err
    }).
    Chdir("share/chimera/client").
    Exec("git", "pull").
    Chdir("../../../").
    Result(func(r interface{}) {
    })

  Task("updateArchetypes").
    Exists("share/chimera/archetypes").
    Catch(func(err error) error {
      cmd := exec.Command("git", "clone", "https://github.com/chimera-rpg/archetypes", "share/chimera/archetypes")
      err = cmd.Run()
      return err
    }).
    Chdir("share/chimera/archetypes").
    Exec("git", "pull").
    Chdir("../../../").
    Result(func(r interface{}) {
    })

  Task("updateMaps").
    Exists("share/chimera/maps").
    Catch(func(err error) error {
      cmd := exec.Command("git", "clone", "https://github.com/chimera-rpg/maps", "share/chimera/maps")
      err = cmd.Run()
      return err
    }).
    Chdir("share/chimera/maps").
    Exec("git", "pull").
    Chdir("../../../").
    Result(func(r interface{}) {
    })

  Task("buildClient").
    Chdir("src/go-client").
    Exec("go", "build", "-v", "-o", "../../bin/client"+exe).
    Catch(func(err error) error {
      fmt.Println(err.Error())
      return nil
    }).
    Chdir("../../").
    Result(func(r interface{}) {
    })

  Task("buildEditor").
    Chdir("src/go-editor").
    Exec("go", "build", "-v", "-o", "../../bin/editor"+exe).
    Catch(func(err error) error {
      fmt.Println(err.Error())
      return nil
    }).
    Chdir("../../").
    Result(func(r interface{}) {
    })

  Task("buildServer").
    Chdir("src/go-server").
    Exec("go", "build", "-v", "-o", "../../bin/server"+exe).
    Catch(func(err error) error {
      fmt.Println(err.Error())
      return nil
    }).
    Chdir("../../").
    Result(func(r interface{}) {
    })

  Task("updateAll").
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
    Run("buildServer")

  Task("watchClient").
    Watch("src/go-client/*.go", "src/go-client/*/*.go").
    Run("buildClient")

  Task("watchEditor").
    Watch("src/go-editor/*.go", "src/go-editor/*/*.go", "src/go-editor/*/*/*.go").
    Run("buildEditor")

	Go()
}
