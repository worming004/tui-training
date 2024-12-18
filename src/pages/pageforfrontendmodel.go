package pages

import (
	"log"

	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/huh/v2"
)

var pageForFrontendClient tea.Model = PageForFrontendModel{}

type PageForFrontendModel struct {
	form *huh.Form
}

func NewPageForFrontendModel() PageForFrontendModel {
	return PageForFrontendModel{
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

func (p PageForFrontendModel) Init() (tea.Model, tea.Cmd) {
	_, cmd := p.form.Init()
	return p, cmd
}

func (p PageForFrontendModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Printf("[PageForFrontendModel] msg: %v, %T", msg, msg)

	form, cmd := p.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		p.form = f
		cmd = cmd
	}

	if p.form.State == huh.StateCompleted {
		m, subCmd := NewDefaultWrapper(NewMainModel()).Init()
		return m, tea.Batch(cmd, subCmd)
	}

	return p, cmd
}

func (p PageForFrontendModel) View() string {
	return p.form.View()
}
