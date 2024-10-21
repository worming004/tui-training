package main

import (
	"os"

	"github.com/charmbracelet/huh"
)

type Runner interface {
	Run() (Runner, error)
}

func main() {
	var runner Runner
	runner = NewApp()
	var err error
	for {
		runner, err = runner.Run()
		if err != nil {
			panic(err)
		}
		if runner == nil {
			os.Exit(0)
		}
	}
}

type ProjectType string

var BackendProjectType = ProjectType("Backend")
var FrontendProjectType = ProjectType("Frontend")
var ClientProjectType = ProjectType("Client")

func ProjectTypeToOption(pt ProjectType) huh.Option[ProjectType] {
	return huh.NewOption(string(pt), pt)
}
