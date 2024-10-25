package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/worming004/tui-training/src/pages"
)

var logFilePath *string = flag.String("logpath", "", "file path for log")

func main() {
	flag.Parse()
	if logFilePath != nil && *logFilePath != "" {
		logFile, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		log.Default().SetOutput(logFile)
		if err != nil {
			panic(err)
		}
	} else {
		log.Default().SetOutput(io.Discard)
	}

	app := pages.NewDefaultWrapper(pages.NewMainModel())
	p := tea.NewProgram(app)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
