package main

import "testing"

func TestStatus_String(t *testing.T) {
	tests := []struct {
		status   Status
		expected string
	}{
		{Installed, "installed"},
		{Outdated, "outdated"},
		{NotInstalled, "not installed"},
		{Status(99), "unknown"},
	}

	for _, tt := range tests {
		if got := tt.status.String(); got != tt.expected {
			t.Errorf("Status(%v).String() = %v, want %v", tt.status, got, tt.expected)
		}
	}
}

func TestStatus_Emoji(t *testing.T) {
	tests := []struct {
		status   Status
		expected string
	}{
		{Installed, "✅"},
		{Outdated, "⚠️"},
		{NotInstalled, "❌"},
		{Status(99), "❓"},
	}

	for _, tt := range tests {
		if got := tt.status.Emoji(); got != tt.expected {
			t.Errorf("Status(%v).Emoji() = %v, want %v", tt.status, got, tt.expected)
		}
	}
}
