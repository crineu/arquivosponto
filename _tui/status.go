package main

// Status represents the stow installation state of a package.
type Status int

const (
	NotInstalled Status = iota
	Outdated
	Installed
)

func (s Status) String() string {
	switch s {
	case Installed:
		return "installed"
	case Outdated:
		return "outdated"
	case NotInstalled:
		return "not installed"
	default:
		return "unknown"
	}
}

func (s Status) Emoji() string {
	switch s {
	case Installed:
		return "✅"
	case Outdated:
		return "⚠️"
	case NotInstalled:
		return "❌"
	default:
		return "❓"
	}
}
