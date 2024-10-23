package pages

import (
	"log"

	tea "github.com/charmbracelet/bubbletea/v2"
)

var quit tea.Model = quitWrapper{}

type quitWrapper struct {
	inner tea.Model
}

func WrapWithQuit(inner tea.Model) quitWrapper {
	return quitWrapper{inner}
}

func NewDefaultWrapper(inner tea.Model) tea.Model {
	return WrapWithQuit(inner)
}

func (q quitWrapper) Init() (tea.Model, tea.Cmd) {
	_, cmd := q.inner.Init()
	return q, cmd
}

func (q quitWrapper) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Printf("[quitWrapper] msg: %v, %T", msg, msg)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return q, tea.Quit
		}
	}

	log.Printf("[quitWrapper] start inner")
	log.Printf("[quitWrapper] inner type %T", q.inner)
	newModel, cmd := q.inner.Update(msg)
	log.Printf("[quitWrapper] stop inner")
	if newModel != q.inner {
		return newModel, cmd
	}
	return q, cmd
}

func (q quitWrapper) View() string {
	return q.inner.View()
}
