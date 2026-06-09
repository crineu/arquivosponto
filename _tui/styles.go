package main

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

// Charm-inspired color palette (simplified)
const (
	ColorPrimary   = "#22C36B" // Green - success, installed, accents
	ColorWarning   = "#FF9F1C" // Orange - outdated, status
	ColorMuted     = "#5C5C5C" // Dark gray - not installed
	ColorSubtle    = "#A0A0A0" // Light gray - help text
	ColorTitleFg   = "#F0F0F0" // Near white for titles
	ColorTitleBg   = "#1A5C3A" // Dark green background for title
)

var (
	// Title
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorTitleFg)).
			Background(lipgloss.Color(ColorTitleBg)).
			Align(lipgloss.Center).
			Padding(0, 2)

	// List items
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().
				PaddingLeft(2).
				Background(lipgloss.Color(ColorPrimary)).
				Foreground(lipgloss.Color("#000000"))

	// Status indicators
	installedStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color(ColorPrimary))
	outdatedStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color(ColorWarning))
	notInstalledStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(ColorMuted))

	// UI elements
	paginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle       = list.DefaultStyles().HelpStyle.
			PaddingLeft(4).
			PaddingBottom(1).
			Foreground(lipgloss.Color(ColorSubtle))

	keyHelpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorSubtle)).
			PaddingLeft(2)

	quitTextStyle = lipgloss.NewStyle().
			Margin(1, 0, 2, 4).
			Foreground(lipgloss.Color(ColorSubtle))

	statusTextStyle = lipgloss.NewStyle().
			Margin(1, 0, 2, 4).
			Foreground(lipgloss.Color(ColorWarning))

	stowTextStyle = lipgloss.NewStyle().
			Margin(1, 0, 1, 4).
			Foreground(lipgloss.Color(ColorSubtle))

	borderStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(ColorPrimary)).
			PaddingRight(2)
)
