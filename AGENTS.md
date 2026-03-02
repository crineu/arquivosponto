# AGENTS.md

## Overview

Personal dotfiles repository managed with **GNU Stow**. Each top-level directory (e.g., \`zsh/\`, \`gitconfig/\`) is a Stow package. Run \`stow &lt;package&gt;\` to symlink configs to \`~/$HOME\`.

Includes:
- Zsh + Prezto customizations
- Mise (asdf successor) for tools
- Configs for Wezterm, Zellij, Neovim, Vim, VSCode, Git, Ruby
- Small **Go TUI** in \`_tui/\` for interactive Stow management (dry-runs)

Repo in Portuguese (PT-BR). See \`README.md\` and \`ATALHOS.md\` for user setup.

## Directory Structure

```
arquivosponto/
├── [stow packages]/    # e.g., zsh/, gitconfig/, wezterm/, zellij/, vscode/, neovim/, mise/, ruby/
├── _tui/               # Go TUI: main.go, stow.go, etc.
├── prezto/             # Vendored Zsh Prezto modules (used by zsh/)
├── scripts/            # Empty (user scripts symlinked to ~/bin)
├── tui/                # Empty
├── README.md           # Setup guide
├── ATALHOS.md          # Quick install (PT)
└── .stowrc             # Stow: --dir=., --target=${HOME}, --dotfiles, --verbose=2
```

## Essential Commands

### Stow (Symlink Management)

```bash
stow zsh                          # Install package symlinks to $HOME
stow -D zsh                       # Uninstall (delete symlinks)
stow -n -v=2 zsh                  # Dry-run: simulate + verbose output
stow *                            # All packages (careful!)
for dir in */; do stow \"$dir\"; done  # From README
```

**TUI for Stow** (recommended):
```bash
cd _tui
go mod tidy
go run main.go
```
- Lists packages (e.g., zsh, gitconfig).
- **Keys**: ↓↑ select, **Enter/i**=stow dry-add, **s**=stow, **d**=delete dry, **r**=restow dry, **c**=check, **q**=quit.

### Mise (Tool Versions)

Config: \`mise/.config/mise/config.toml\`

```bash
stow mise
mise use --global go@1.22.4          # Or latest for lazygit, ripgrep, bat, neovim, zellij, etc.
mise ls                              # List installed
mise current                         # Active versions
```

### Development (Go _tui/)

```bash
cd _tui
go mod tidy                  # Fix imports (lsp errors otherwise)
go build ./cmd/...           # Build
go run main.go               # Run TUI
go test ./...                # Tests (none yet)
golangci-lint run            # Lint (mise tool)
```

### Zsh/Prezto

```bash
stow zsh
chsh -s /usr/bin/zsh
```
- Extras: \`*.zsh\` in \`~/.zsh_extras/\` auto-sourced.
- Completions: \`_* \` files auto-used.

## Code Patterns

- **Go (_tui/)**: Bubble Tea TUI (Charmbracelet). Files: \`main.go\` (list+keys), \`stow.go\` (stow exec), \`home_finder.go\`, \`folders.go\`.
  - Keymap: Custom bindings (s/d/r/c).
  - Styles: Lipgloss themes (greens/pinks).

- **Configs**:
  | Format | Examples |
  |--------|----------|
  | TOML   | mise/config.toml |
  | Lua    | wezterm.lua |
  | KDL    | zellij/config.kdl |
  | JSON   | VSCode settings/keybindings |
  | INI    | gitconfig |

- **Zsh**: Prezto modules + overrides in \`zsh/\`. Git aliases/prompt, Ruby completions.

## Testing & Linting

- **Go**: `go test ./...` (no tests yet). `golangci-lint run`.
- No other test suites or CI configs found.
- LSP: Go errors until `go mod tidy`.

## Gotchas

- **Stow Target**: Always `$HOME` (\`.stowrc\` with \`--dotfiles\`).
- **Dotfiles**: Files named \`dot-zshrc\` link as \`.zshrc\` in HOME.
- **Go Imports**: Run `go mod tidy` first (missing deps cause LSP fails).
- **Prezto Vendored**: Modules in \`prezto/\`; customize via \`zsh/\`.
- **Scripts Empty**: `~/bin/` gets symlinks when populated.
- **Mise Ruby**: `idiomatic_version_file_enable_tools = [\"ruby\"]` (reads .ruby-version).
- Changes: Test stow dry-run before applying. Commit with `gitconfig` aliases (gba, gst, etc.).

## Memory for Agents

- Build: `go build` (_tui/)
- Test: `go test ./...`
- Lint: `golangci-lint run`
- Run: `go run main.go` or `stow <pkg>`

**Only document observed facts. Updated from codebase scan.**

## Recommended Improvements (Linux-only)

1. Substitua Prezto por zplug/zinit (leve, auto-update). Migração: `zplug "sorin-ionescu/prezto"`, porte customs para zsh_extras.
2. Adicione: `starship` prompt (cross-shell); `direnv` para envs; `fzf` fuzzy.
3. CI/GitHub Actions: Test stow dry-run + lint configs.

