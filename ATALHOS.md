# Recomendado: Install mise, wezterm and zsh

    sudo apt install stow
    sudo pacman -Sy stow

# Change shell to ZSH
    stow  zsh
    chsh -s /usr/bin/zsh

# install other tools with mise
    stow mise

    mise use --global git@latest
    # sudo apt install make libssl-dev libghc-zlib-dev libcurl4-gnutls-dev libexpat1-dev gettext unzip (prereq for debian)
    mise use --global zellij@latest
    mise use --global tmux@latest
    mise use --global neovim@latest
    mise use --global lazygit@latest

    stow wezterm
    stow gitconfig
    stow zellij
    stow neovim
    stow vscode
