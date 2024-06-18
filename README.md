Meus arquivos ponto – famosos _dotfiles_

# :rocket: Instalação

Pré-requisitos:
* [Git](https://git-scm.com/)
* [GNU Stow](https://www.gnu.org/software/stow/)

```bash
git clone git@github.com:crineu/arquivosponto.git ~/arquivosponto
git clone https://github.com/crineu/arquivosponto ~/arquivosponto
cd ~/arquivosponto
stow [nome_do_diretório]
```

---

# Lista de ferramentas recomendadas

## :hammer: Mise-en-place https://mise.jdx.dev/

```bash
curl https://mise.run | sh
stow mise
mise use --global [plugin]@[version]
mise user --global [ruby zellij lazygit neovim ripgrep bat]@latest
```

## :hammer: WezTerm https://wezfurlong.org/wezterm/


```bash
stow wezterm
```

* `~/.wezterm.lua` # configurações, atalhos etc


## :hammer: ZSH

* Install [Z Shell](http://zsh.sourceforge.net/Doc/Release/Introduction.html)

```bash
stow zsh
chsh -s /usr/bin/zsh
```

Define o _shell_ como ZSH e habilita:

* Prompt com git status na linha de comando e alias:
* ZMV
  - `zmv  'batata-(*)' 'banana.$1'` <-- :scream:

### Ordem de inicialização dos arquivos:

* `.zshenv`: invocado sempre; deve ser conciso e meramente inicializar as variáveis necessárias;
* `.zlogin`: carregado em _login shells_ após `.zshrc`; compila _zcompdump_ em segundo plano pois é processo lento e feito 1x / login;
* `.zprofile` : similar a `.zlogin` mas carregado antes de `.zshrc` (`.zprofile` e  `.zshrc` são ignorados em _non-login non-interactive shells - so I learned a trick from Prezto that declares environment variables in _.zprofile_ and uses .zshenv_ to source _.zprofile_ (e.g. .zprofile and .zshenv). This way, non-login non-interactive shells will receive proper variable initialisations).
* `.zshrc`: carregado em _interactive shells_; contém a parte principal das configurações do ZSH.

### Personalizações

Arquivos `.zsh` em `$HOME/.zsh_extras` serão automaticamente _sourced_ quando o login for efetuado.



## :hammer: Scripts

```bash
stow scripts
```

Scripts úteis ligados em `~/bin`


## :hammer: Zellij

```bash
stow zellij
```

* `~/.config/zellij/config/kdl` configurações do Zellij


## :hammer: Gitconfig

```bash
stow gitconfig
git config -l --show-origin     # mostra configurações e origem
```

* `~/.gitconfig` configurações, definição de arquivos padrão e aliases:
  - `gba`, `gbb`: branchs e conexões remotas;
  - `gst`, `gsc`: git status, git commit
  - `glg`, `gld`, `glb`, `glc` - distintos tipos de log
* `~/.gitconfig.user` crie esse arquivo para configurações pessoais
* `~/.gitconfig.commit-template.txt` como o nome indica
* `~/.gitignore` para ignores file



## :hammer: VS Code

```bash
stow vscode
```

* `~/.config/Code - OSS/User/settings.json`
* `~/.config/Code - OSS/User/keybindings.json`



---
---

# :volcano: Remoção completa

```bash
cd ~/arquivosponto

for folder in */ ; do
  stow -D -n "$folder"
done

cd ..
rm -rf arquivosponto
```


---

# Ferramentas Legadas

## :hammer: ASDF https://asdf-vm.com/#/core-manage-asdf-vm

> Substituído pelo Mise

```bash
stow asdf
asdf plugin-list-all
asdf plugin-add {ruby, go, zellij, tmux}
asdf install {ruby, go, zellij, tmux} _version_number
asdf global {ruby, go, zellij, tmux} _version_number
```

## :hammer: Alacritty

> Substituído pelo WezTerm

```bash
stow alacritty
```

* `~/.alacritty` configurações padrão do terminal Alacritty


## :hammer: Tmux

> Substituído pelo Zellij

```bash
stow tmux
```

* `~/.tmux.conf` configurações do tmux
* `~/.tmux.conf.user` crie esse arquivo para configurações pessoais
* `~/bin/tmux_dev_env.sh` exemplo de automatização ao iniciar tmux (necessita `~/bin` presente no `$PATH`)
* `CTRL` + `a` como comando padrão


## :hammer: Vim

> Substituído pelo NeoVIM

```bash
stow vim
# vim versão > 8
git clone https://github.com/preservim/nerdtree.git ~/.vim/pack/vendor/start/nerdtree
vim -u NONE -c "helptags ~/.vim/pack/vendor/start/nerdtree/doc" -c q
```

* `~/.vimrc` configurações do vim
* `~/.vimrc.before` crie esse arquivo para configurações pessoais
* plugins recomendado: [NERDTree](https://github.com/preservim/nerdtree) (`CTRL` + `n`)
