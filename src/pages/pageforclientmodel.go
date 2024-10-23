package pages

import (
	"log"

	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/huh/v2"
)

var clientPageModel tea.Model = PageForClientModel{}

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
	_, cmd := p.form.Init()
	return p, cmd
}

// Update implements tea.Model.
func (p PageForClientModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Printf("[PageForClientModel] msg: %v, %T", msg, msg)

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

// View implements tea.Model.
func (p PageForClientModel) View() string {
	return p.form.View()
}
