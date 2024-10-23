package pages

import (
	"log"

	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/huh/v2"
)

var pageForBackendClient tea.Model = PageForBackendModel{}

type PageForBackendModel struct {
	form *huh.Form
}

func NewPageForBackendModel() PageForBackendModel {
	return PageForBackendModel{
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

func (p PageForBackendModel) Init() (tea.Model, tea.Cmd) {
	_, cmd := p.form.Init()
	return p, cmd
}

func (p PageForBackendModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Printf("[PageForBackendModel] msg: %v, %T", msg, msg)
	form, cmd := p.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		p.form = f
	}

	if p.form.State == huh.StateCompleted {
		m, subCmd := NewDefaultWrapper(NewMainModel()).Init()
		return m, tea.Batch(cmd, subCmd)
	}

	return p, cmd
}

func (p PageForBackendModel) View() string {
	return p.form.View()
}
