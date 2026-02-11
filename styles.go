package main

import (
	"github.com/charmbracelet/lipgloss"
)

// UI Styles
var (
	// Reduced horizontal padding to 0 to minimize sidebar margins as requested
	appStyle = lipgloss.NewStyle().Padding(1, 0, 0, 0)

	// Header Styles
	logoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("212")).
			MarginRight(1).
			SetString("\uf489") //

	channelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Bold(true).
			MarginRight(1).
			SetString("#general")

	dividerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			MarginRight(1).
			SetString("|")

	topicStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("243")). // Grey
			MarginRight(1).                    // Reduced margin to fit new divider
			SetString("TOPIC: Discussion")

	searchBaseStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Padding(0, 1)

	iconBoxStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Padding(0, 1).
			MarginLeft(1).
			Align(lipgloss.Center)

	// The big wrapper for everything
	headerContainerStyle = lipgloss.NewStyle().
				Border(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("240")).
				Padding(0, 1).
				MarginTop(1)

	// Status Line Style (Re-added)
	statusLineStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("212")).
			Padding(0, 1)

	// Main Content Area Style (Empty Border Box)
	mainContentStyle = lipgloss.NewStyle().
				Border(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("240")).
				Padding(0, 1)

	// Message Input Box Style
	messageBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("212")). // Pink border
			Padding(0, 1).
			MarginTop(0)

	// Sidebar Styles
	// Added MarginTop(1) to align with Header
	leftSidebarStyle = lipgloss.NewStyle().
				Border(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("240")).
				Padding(0, 1).
				MarginTop(1)

	rightSidebarStyle = lipgloss.NewStyle().
				Border(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("240")).
				Padding(0, 1).
				MarginTop(1)
)
