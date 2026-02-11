package main

import (
	"github.com/charmbracelet/lipgloss"
)

func (m *model) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	// Sidebar Widths
	// Fixed content width of 20
	leftSidebarContentWidth := 20
	rightSidebarContentWidth := 20

	// --- 1. HEADER ---
	leftSide := lipgloss.JoinHorizontal(lipgloss.Center,
		logoStyle.String(),
		channelStyle.String(),
		dividerStyle.String(),
		topicStyle.String(),
		dividerStyle.String(),
	)
	leftWidth := lipgloss.Width(leftSide)

	bellIcon := iconBoxStyle.Render("\uf0f3") //
	infoIcon := iconBoxStyle.Render("\uf05a") //
	rightSide := lipgloss.JoinHorizontal(lipgloss.Center, bellIcon, infoIcon)
	rightWidth := lipgloss.Width(rightSide)

	// Header Container Width calculation
	// headerContainerStyle has border(2) + padding(2) = 4 extra width
	headerContentWidth := m.centerRenderedWidth - 4

	// Search Bar Width Calculation
	targetTotalWidth := headerContentWidth - leftWidth - rightWidth
	// searchBaseStyle adds: 2 (padding L/R)
	searchContentWidth := targetTotalWidth - 2
	if searchContentWidth < 10 {
		searchContentWidth = 10
	}

	// Dynamic style for search input
	var searchInputView string
	if m.textInput.Focused() {
		pinkStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
		m.textInput.TextStyle = pinkStyle
		m.textInput.PromptStyle = pinkStyle
		searchInputView = searchBaseStyle.Width(searchContentWidth).Render(m.textInput.View())
	} else {
		grayStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
		m.textInput.TextStyle = grayStyle
		m.textInput.PromptStyle = grayStyle
		searchInputView = searchBaseStyle.Width(searchContentWidth).Render(m.textInput.View())
	}

	headerContent := lipgloss.JoinHorizontal(lipgloss.Center, leftSide, searchInputView, rightSide)

	// Set width on container to ensure it fills space
	header := headerContainerStyle.Width(headerContentWidth).Render(headerContent)

	// --- 2. STATUS LINE ---
	// Status line has no border. Bordered elements render at centerRenderedWidth-2,
	// so subtract 2 to align with them.
	statusLine := statusLineStyle.
		Width(m.centerRenderedWidth - 2).
		Render("MESSAGE-BUFFER")

	// --- 3. BOTTOM MESSAGE INPUT ---
	promptColor := lipgloss.Color("240")
	borderColor := lipgloss.Color("240")
	if m.messageInput.Focused() {
		promptColor = lipgloss.Color("212")
		borderColor = lipgloss.Color("212")
	}

	prompt := lipgloss.NewStyle().Foreground(promptColor).Render("> ")
	icons := lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render(" \uee49 \U000F0066")

	// Input Box Width Logic
	// messageBox has border(2) + padding(2) = 4 extra width
	messageBoxContentWidth := m.centerRenderedWidth - 4

	inputContent := lipgloss.JoinHorizontal(lipgloss.Top,
		prompt,
		m.messageInput.View(),
		icons,
	)

	currentMessageBoxStyle := messageBoxStyle.BorderForeground(borderColor)
	messageBox := currentMessageBoxStyle.
		Width(messageBoxContentWidth). // Sets content width
		Render(inputContent)

	// --- 4. MAIN CONTENT (Border Box) ---
	headerH := lipgloss.Height(header)
	statusH := lipgloss.Height(statusLine)
	messageH := lipgloss.Height(messageBox)

	// Total Available Height Calculation
	// m.height - appPadding Top(1)
	availableMainHeight := m.height - 1 - headerH - statusH - messageH
	if availableMainHeight < 0 {
		availableMainHeight = 0
	}

	mainContent := mainContentStyle.
		Width(m.centerRenderedWidth - 4).
		Height(availableMainHeight).
		Render("")

	// Compose Center Column
	centerColumn := lipgloss.JoinVertical(lipgloss.Left,
		header,
		statusLine,
		mainContent,
		messageBox,
	)

	// --- 5. SIDEBARS ---
	// Calculate sidebar height to match the center column exactly.
	// centerColumn height includes all vertical components + their margins.
	// Sidebars have MarginTop(1) and Border(2).
	// So Content Height = CenterHeight - Margin(1) - Border(2) = CenterHeight - 3.
	sidebarHeight := lipgloss.Height(centerColumn)
	sidebarContentHeight := sidebarHeight - 3

	if sidebarContentHeight < 0 {
		sidebarContentHeight = 0
	}

	leftSidebar := leftSidebarStyle.
		Width(leftSidebarContentWidth).
		Height(sidebarContentHeight).
		Render("")

	rightSidebar := rightSidebarStyle.
		Width(rightSidebarContentWidth).
		Height(sidebarContentHeight).
		Render("")

	// --- 6. COMBINE COLUMNS ---
	finalView := lipgloss.JoinHorizontal(lipgloss.Top,
		leftSidebar,
		centerColumn,
		rightSidebar,
	)

	return appStyle.Render(finalView)
}
