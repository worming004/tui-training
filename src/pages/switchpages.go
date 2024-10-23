package pages

import tea "github.com/charmbracelet/bubbletea/v2"

type SwitchPageMsg int

const (
	SwitchToClientPageMsg   SwitchPageMsg = iota
	SwitchToBackendPageMsg  SwitchPageMsg = iota
	SwitchToFrontendPageMsg SwitchPageMsg = iota
	SwitchToMainPageMsg     SwitchPageMsg = iota
)

func MsgToCmd(msg tea.Msg) tea.Cmd {
	return func() tea.Msg { return msg }
}
