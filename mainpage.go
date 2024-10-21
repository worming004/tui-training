package main

import "github.com/charmbracelet/huh"

// Represent responses from user
type MainAppResponses struct {
	ProjectType
	SubPageResponses interface{}
}
type App struct {
	*MainAppResponses
}

func NewApp() *App {
	return &App{
		MainAppResponses: &MainAppResponses{},
	}
}

func (a *App) Run() (Runner, error) {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[ProjectType]().
				Description("Type de projet a construire").
				Title("Type de projet").
				Options(
					ProjectTypeToOption(BackendProjectType),
					ProjectTypeToOption(FrontendProjectType),
					ProjectTypeToOption(ClientProjectType),
				).
				Value(&a.MainAppResponses.ProjectType),
		),
	)

	err := form.Run()

	if err != nil {
		return nil, err
	}

	switch a.MainAppResponses.ProjectType {
	case ClientProjectType:
		return NewPageForClient(), nil
	}

	return nil, nil
}
