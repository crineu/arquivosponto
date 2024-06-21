#
# Executes commands at the start of an interactive session.
#
# Authors:
#   Sorin Ionescu <sorin.ionescu@gmail.com>
#

# Custom folder where personal configuration must be added
if [ -d $HOME/.zsh_extras/ ]; then
  if [ "$(ls -A $HOME/.zsh_extras/)" ]; then
    for config_file ($HOME/.zsh_extras/*.zsh) source $config_file
  fi
fi

# Completions files in the same folder; must start with "_"
fpath=(~/.zsh_extras $fpath)
# initialise completions with ZSH's compinit
autoload -Uz compinit && compinit


# Source Prezto.
if [[ -s "${ZDOTDIR:-$HOME}/arquivosponto/prezto/init.zsh" ]]; then
  source "${ZDOTDIR:-$HOME}/arquivosponto/prezto/init.zsh"
fi

# Customize to your needs...

# No shared historby between sessions
unsetopt SHARE_HISTORY
setopt NO_SHARE_HISTORY
