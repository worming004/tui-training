package main

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

type PageForClientResponses struct {
	ProjectName string
}
type PageForClient struct {
	*PageForClientResponses
}

func (p *PageForClient) Run() (Runner, error) {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("ProjectName").
				Value(&p.PageForClientResponses.ProjectName),
		),
	)

	err := form.Run()
	if err != nil {
		return nil, err
	}

	return &CreateClientApp{p.PageForClientResponses}, nil
}

func NewPageForClient() *PageForClient {
	return &PageForClient{
		PageForClientResponses: &PageForClientResponses{},
	}
}

type CreateClientApp struct {
	*PageForClientResponses
}

func (c *CreateClientApp) Run() (Runner, error) {
	fmt.Printf("Creating client app with name %s\n", c.PageForClientResponses.ProjectName)
	return nil, nil
}
