package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/huh/v2"
)

var page tea.Model = PageForClientModel{}

type PageForClientModel struct {
	form *huh.Form
}

func NewPageForClientModel() PageForClientModel {
	return PageForClientModel{
		form: huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("ProjectName").
					Description("Nom du projet").
					Key("ProjectName"),
			),
		),
	}
}

// Init implements tea.Model.
func (p PageForClientModel) Init() (tea.Model, tea.Cmd) {
	log.Println("PageForClient Init")
	_, cmd := p.form.Init()
	return p, cmd
}

// Update implements tea.Model.
func (p PageForClientModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Printf("Received message from client page: %+v, %T", msg, msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return p, tea.Quit
		}
	}

	cmds := []tea.Cmd{}
	form, cmd := p.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		p.form = f
		cmds = append(cmds, cmd)
	}

	if p.form.State == huh.StateCompleted {
		pn := p.form.GetString("ProjectName")
		log.Printf("Project name: %s\n", pn)
		return p, tea.Quit
	}

	return p, tea.Batch(cmds...)
}

// View implements tea.Model.
func (p PageForClientModel) View() string {
	return p.form.View()
}
