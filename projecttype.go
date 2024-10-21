package main

import "github.com/charmbracelet/huh/v2"

type ProjectType string

var BackendProjectType = ProjectType("Backend")
var FrontendProjectType = ProjectType("Frontend")
var ClientProjectType = ProjectType("Client")

func ProjectTypeToOption(pt ProjectType) huh.Option[ProjectType] {
	return huh.NewOption(string(pt), pt)
}
