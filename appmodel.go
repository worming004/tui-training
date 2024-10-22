package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea/v2"
)

var appmodel tea.Model = AppModel{}

// AppModel is the top-level model for the application.
// It handle quit with ctrl+c and model switching.
type AppModel struct {
	activeModel tea.Model
}

func NewDefaultAppModel() AppModel {
	return AppModel{
		activeModel: NewMainModel(),
	}
}

// Init implements tea.Model.
func (a AppModel) Init() (tea.Model, tea.Cmd) {
	_, cmd := a.activeModel.Init()
	return a, cmd
}

// Update implements tea.Model.
func (a AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Printf("msg: %v, with type %T", msg, msg)
	cmds := []tea.Cmd{}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return a, tea.Quit
		}

	case SwitchPageMsg:
		log.Printf("SwitchPageMsg: %v", msg)
		switch msg {
		case SwitchToClientPageMsg:
			a.activeModel = NewPageForClientModel()
			_, cmd := a.activeModel.Init()
			return a, cmd

		case SwitchToBackendPageMsg:
			a.activeModel = NewPageForBackendModel()
			_, cmd := a.activeModel.Init()
			return a, cmd

		case SwitchToFrontendPageMsg:
			a.activeModel = NewPageForFrontendModel()
			_, cmd := a.activeModel.Init()
			return a, cmd

		case SwitchToMainPageMsg:
			a.activeModel = NewDefaultAppModel()
			_, cmd := a.activeModel.Init()
			return a, cmd
		}
	}

	log.Printf("didn't match any switch case")

	_, cmd := a.activeModel.Update(msg)
	cmds = append(cmds, cmd)
	return a, tea.Batch(cmds...)
}

// View implements tea.Model.
func (a AppModel) View() string {
	return a.activeModel.View()
}
