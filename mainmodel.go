package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/huh/v2"
)

// Just to ensure AppModel is a tea.Model
var model tea.Model = MainModel{}

// Represent responses from user
type MainAppResponses struct {
	SubPageResponses interface{}
}
type MainModel struct {
	form *huh.Form
	*MainAppResponses
}

func NewMainModel() MainModel {
	return MainModel{
		MainAppResponses: &MainAppResponses{},
		form: huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[ProjectType]().
					Title("Type de projet").
					Description("Type de projet a construire").
					Key("ProjectType").
					Options(
						ProjectTypeToOption(BackendProjectType),
						ProjectTypeToOption(FrontendProjectType),
						ProjectTypeToOption(ClientProjectType),
					),
			),
		),
	}
}

// Init implements tea.Model.
func (a MainModel) Init() (tea.Model, tea.Cmd) {
	log.Println("AppModel Init")
	_, cmd := a.form.Init()
	return a, cmd
}

// Update implements tea.Model.
func (a MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Printf("Received message from main page: %+v, %T", msg, msg)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return a, tea.Quit
		}
	}

	cmds := []tea.Cmd{}
	form, cmd := a.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		a.form = f
		cmds = append(cmds, cmd)
	}

	if a.form.State == huh.StateCompleted {
		pt := a.form.Get("ProjectType").(ProjectType)
		switch pt {
		case ClientProjectType:
			client := NewPageForClientModel()
			// Doesn't not work without Init explicitely. Internally, it is using another huh form that require init
			_, cmd = client.Init()
			// Am I supposed to return a new model ? MainModel should handle the whole app and work as decorator for each sub-part of app ?
			return client, tea.Batch(append(cmds, cmd)...)

		case BackendProjectType, FrontendProjectType:
			return a, tea.Batch(cmds...)
		}
		return a, tea.Quit
	}

	return a, cmd
}

// View implements tea.Model.
func (a MainModel) View() string {
	return a.form.View()
}
