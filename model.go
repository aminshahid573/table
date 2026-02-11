package main

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	width               int
	height              int
	centerRenderedWidth int
	inputWidth          int             // textarea character width, for visual line wrapping
	textInput           textinput.Model // Search bar
	messageInput        textarea.Model  // Message input
}

func initialModel() model {
	// Search Input
	ti := textinput.New()
	ti.Placeholder = "Search"
	ti.Prompt = "\uf002 " //
	ti.CharLimit = 156
	ti.Width = 20
	// Style for search input
	color240 := lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	ti.PromptStyle = color240
	ti.PlaceholderStyle = color240
	ti.TextStyle = color240
	ti.Cursor.Style = color240

	// Message Input (Textarea)
	ta := textarea.New()
	ta.Placeholder = "Type a Message or command (use / for actions)"
	ta.ShowLineNumbers = false
	ta.SetHeight(1)
	ta.Prompt = ""
	ta.Focus() // Focus message input by default

	// Clear default styles to remove "light white-grey focus"
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()
	ta.FocusedStyle.Base = lipgloss.NewStyle()
	ta.BlurredStyle.CursorLine = lipgloss.NewStyle()
	ta.BlurredStyle.Base = lipgloss.NewStyle()

	return model{
		textInput:    ti,
		messageInput: ta,
	}
}

func (m *model) recalcLayout() {
	leftSidebarRenderedWidth := 22  // 20 content + 2 border
	rightSidebarRenderedWidth := 22 // 20 content + 2 border

	m.centerRenderedWidth = m.width - leftSidebarRenderedWidth - rightSidebarRenderedWidth
	if m.centerRenderedWidth < 40 {
		m.centerRenderedWidth = 40
	}

	// Textarea width: same formula as View's messageBox layout
	// messageBoxStyle has border(2) + padding(2) = 4 extra
	messageBoxContentWidth := m.centerRenderedWidth - 4
	prompt := "> "
	icons := " \uee49 \U000F0066"
	promptW := lipgloss.Width(prompt)
	iconsW := lipgloss.Width(icons)
	inputWidth := messageBoxContentWidth - promptW - iconsW - 2
	if inputWidth < 1 {
		inputWidth = 1
	}
	m.inputWidth = inputWidth
	m.messageInput.SetWidth(inputWidth)

	// Textinput width: same formula as View's header layout
	headerContentWidth := m.centerRenderedWidth - 4
	leftSide := lipgloss.JoinHorizontal(lipgloss.Center,
		logoStyle.String(),
		channelStyle.String(),
		dividerStyle.String(),
		topicStyle.String(),
		dividerStyle.String(),
	)
	leftWidth := lipgloss.Width(leftSide)

	bellIcon := iconBoxStyle.Render("\uf0f3")
	infoIcon := iconBoxStyle.Render("\uf05a")
	rightSide := lipgloss.JoinHorizontal(lipgloss.Center, bellIcon, infoIcon)
	rightWidth := lipgloss.Width(rightSide)

	targetTotalWidth := headerContentWidth - leftWidth - rightWidth
	searchContentWidth := targetTotalWidth - 2
	if searchContentWidth < 10 {
		searchContentWidth = 10
	}
	m.textInput.Width = searchContentWidth - 2
}

func (m *model) Init() tea.Cmd {
	return tea.Batch(
		tea.SetWindowTitle("Bubble Tea TUI"),
		textinput.Blink,
		textarea.Blink,
	)
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "tab":
			if m.textInput.Focused() {
				m.textInput.Blur()
				m.messageInput.Focus()
			} else {
				m.messageInput.Blur()
				m.textInput.Focus()
			}
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	// Update inputs
	m.textInput, cmd = m.textInput.Update(msg)
	cmds = append(cmds, cmd)
	m.messageInput, cmd = m.messageInput.Update(msg)
	cmds = append(cmds, cmd)

	// Recalculate layout (width must be set before dynamic height check)
	if m.width > 0 {
		m.recalcLayout()
	}

	// Post-update Logic for dynamic height
	// Count visual lines (including wrapping) using Lipgloss's wrapper
	visualLines := 0
	if m.inputWidth > 0 {
		// Use lipgloss to simulate wrapping behavior
		wrapped := lipgloss.NewStyle().Width(m.inputWidth).Render(m.messageInput.Value())
		visualLines = lipgloss.Height(wrapped)
	}

	if visualLines > 1 {
		m.messageInput.SetHeight(2)
	} else {
		m.messageInput.SetHeight(1)
	}
	// Limit to 2 lines max visible
	if m.messageInput.Height() > 2 {
		m.messageInput.SetHeight(2)
	}

	return m, tea.Batch(cmds...)
}
