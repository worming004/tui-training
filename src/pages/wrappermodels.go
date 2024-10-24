package pages

import (
	"log"

	tea "github.com/charmbracelet/bubbletea/v2"
)

var quit tea.Model = quitWrapper{}
var handler tea.Model = dummyCommandHandler{}

type quitWrapper struct {
	inner tea.Model
}

func WrapWithQuit(inner tea.Model) quitWrapper {
	return quitWrapper{inner}
}

func NewDefaultWrapper(inner tea.Model) tea.Model {
	// return WrapWithQuit(WrapWithHandler(inner))
	return WrapWithHandler(WrapWithQuit(inner))
}

func (q quitWrapper) Init() (tea.Model, tea.Cmd) {
	_, cmd := q.inner.Init()
	return q, cmd
}

func (q quitWrapper) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return q, tea.Quit
		}
	}

	newModel, cmd := q.inner.Update(msg)
	if newModel != q.inner {
		return newModel, cmd
	}
	return q, cmd
}

func (q quitWrapper) View() string {
	return q.inner.View()
}

type dummyCommandHandler struct {
	inner tea.Model
}

func WrapWithHandler(inner tea.Model) dummyCommandHandler {
	log.Printf("[DummyCommandHandler] Wrapping with handler")
	return dummyCommandHandler{inner}
}

// Init implements tea.Model.
func (d dummyCommandHandler) Init() (tea.Model, tea.Cmd) {
	_, cmd := d.inner.Init()
	return d, cmd
}

// Update implements tea.Model.
func (d dummyCommandHandler) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Printf("[DummyCommandHandler] %+v, %T", msg, msg)
	switch msg := msg.(type) {
	case SendForm:
		log.Printf("[DummyCommandHandler] Command received, SendForm: %+v", msg)
	}

	newModel, cmd := d.inner.Update(msg)
	if newModel != d.inner {
		log.Printf("[DummyCommandHandler] newModel != d.inner, update")
		return newModel, cmd
	}
	return d, cmd
}

// View implements tea.Model.
func (d dummyCommandHandler) View() string {
	log.Printf("[DummyCommandHandler] View")
	return d.inner.View()
}
