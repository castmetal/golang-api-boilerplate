package ports

import (
	"context"
	"fmt"

	"github.com/castmetal/golang-api-boilerplate/src/config"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

type ServerInterface interface {
	CreateServer(ctx context.Context)
}

func IsDev(config config.EnvStruct) bool {
	isDev := true
	if config.Environment == "prod" {
		isDev = false
	}

	return isDev
}

func ServerLog(config config.EnvStruct) {
	size := 1

	if term.IsTerminal(0) {
		size = 0
	}

	terminalWidth, _, err := term.GetSize(size)
	if err != nil {
		return
	}

	fmt.Print("\n\n")

	headerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("15")).
		Background(lipgloss.Color("#7829E1")).
		Bold(true).
		Width(terminalWidth).
		Align(lipgloss.Left)

	contentStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("7")).
		Background(lipgloss.Color("0")).
		Width(terminalWidth).
		Padding(1)

	footerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("0")).
		Background(lipgloss.Color("#E2E2E2")).
		Bold(true).
		Width(terminalWidth).
		Align(lipgloss.Center)

	header := headerStyle.Render("Castmetal - github.com/castmetal API")
	content := contentStyle.Render(config.ApiName)
	footer := footerStyle.Render("© 2023")

	fmt.Println(header)
	fmt.Println(content)
	fmt.Println(footer)

	fmt.Print("\n\n\n")
}
