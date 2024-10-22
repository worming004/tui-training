package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/huh/v2"
)

// Just to ensure AppModel is a tea.Model
var mainmodel tea.Model = MainPageModel{}

// Represent responses from user
type MainAppResponses struct {
	SubPageResponses interface{}
}
type MainPageModel struct {
	form *huh.Form
	*MainAppResponses
}

func NewMainModel() MainPageModel {
	return MainPageModel{
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
func (m MainPageModel) Init() (tea.Model, tea.Cmd) {
	log.Println("AppModel Init")
	_, cmd := m.form.Init()
	return m, cmd
}

// Update implements tea.Model.
func (m MainPageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := []tea.Cmd{}

	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	if m.form.State == huh.StateCompleted {
		pt := m.form.Get("ProjectType").(ProjectType)
		switch pt {
		case ClientProjectType:
			cmds = append(cmds, MsgToCmd(SwitchToClientPageMsg))
		case BackendProjectType:
			cmds = append(cmds, MsgToCmd(SwitchToBackendPageMsg))
		case FrontendProjectType:
			cmds = append(cmds, MsgToCmd(SwitchToFrontendPageMsg))
		}
	}

	return m, tea.Batch(cmds...)
}

// View implements tea.Model.
func (m MainPageModel) View() string {
	return m.form.View()
}
