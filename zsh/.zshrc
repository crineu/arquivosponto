#
# Executes commands at the start of an interactive session.
#
# Authors:
#   Sorin Ionescu <sorin.ionescu@gmail.com>
#

# Source Prezto.
if [[ -s "${ZDOTDIR:-$HOME}/arquivosponto/prezto/init.zsh" ]]; then
  source "${ZDOTDIR:-$HOME}/arquivosponto/prezto/init.zsh"
fi

# Customize to your needs...

# Show links created
alias stow='stow --verbose'

# No shared historby between sessions
unsetopt SHARE_HISTORY
setopt NO_SHARE_HISTORY