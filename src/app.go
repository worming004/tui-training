package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/worming004/tui-training/src/pages"
)

func main() {
	logFile, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.Default().SetOutput(logFile)

	app := pages.NewDefaultWrapper(pages.NewMainModel())
	p := tea.NewProgram(app)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
