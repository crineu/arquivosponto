Meus arquivos ponto (famosos _dotfiles_)

## :underage: Instalação

Pré-requisitos: [GNU Stow](https://www.gnu.org/software/stow/) e [Z Shell](http://zsh.sourceforge.net/Doc/Release/Introduction.html)

```bash
git clone git@github.com:crineu/arquivosponto.git ~/arquivosponto
cd ~/arquivosponto

stow --verbose zsh
chsh -s /usr/bin/zsh
```

Clona repositório, define o _shell_ como ZSH e habilita:

* Prompt deveras bacana com git status na linha de comando e alias:
    - `gba`, `gbb`: branchs e conexões remotas;
    - `gst`, `gsc`: git status, git commit
    - `glg`, `gld`, `glb`, `glc` - distintos tipos de log
* ZMV!
    - `zmv  'batata-(*)' 'banana.$1'` <-- :scream:


### :earth_asia: Personalizações

Qualquer arquivo `.zsh` em `$HOME/.zsh_extras` será automaticamente _sourced_ quando o login for efetuado.

Exemplo: arquivo `$HOME/.zsh_extras/docker.zsh` para adicionar _aliases_ a comandos docker:

```bash
alias dkls='docker image ls'
```


### :earth_americas: Gitconfig

* `~/.gitconfig` várias configurações do git, e definição de arquivos padrão:
* `~/.gitconfig.user` crie esse arquivo para configurações pessoais
* `~/.gitconfig.commit-template.txt` como o nome indica
* `~/.gitignore` para ignores file

```bash
stow gitconfig
git config -l --show-origin  # mostra configurações e origem
```


### :earth_africa: Tmux

* `~/.tmux.conf` configurações do tmux
* `~/.tmux.conf.user` crie esse arquivo para configurações pessoais
* `CTRL` + `a` como tecla padrão

```bash
stow tmux
```

### :earth_asia: Vim

* `~/.vimrc` configurações do vim
* `~/.vimrc.before` crie esse arquivo para configurações pessoais
* plugins recomendado: [NERDTree](https://github.com/preservim/nerdtree) (`CTRL` + `n`)

```bash
stow vim

# vim versão > 8
git clone https://github.com/preservim/nerdtree.git ~/.vim/pack/vendor/start/nerdtree
vim -u NONE -c "helptags ~/.vim/pack/vendor/start/nerdtree/doc" -c q
```

### :earth_americas: ASDF

Recomenda-se a instalação do [ASDF](https://asdf-vm.com/#/core-manage-asdf-vm). Após:

```bash
stow asdf

asdf plugin-list-all
asdf plugin-add {ruby, rust, docker, tmux}
asdf install {ruby, rust, docker, tmux} _version_number
```

### :earth_africa: VS Code

* `~/.config/Code - OSS/User/settings.json`
* `~/.config/Code - OSS/User/keybindings.json`

```bash
stow vscode
```


---

## :volcano: Remoção completa :volcano:

```bash
cd ~/arquivosponto

for folder in */ ; do
  stow -D -n "$folder"
done

cd ..
rm -rf arquivosponto
```

---

<!-- :new_moon: -->
<!-- :waxing_crescent_moon: -->
<!-- :first_quarter_moon: -->
<!-- :waxing_gibbous_moon: -->
<!-- :full_moon: -->
<!-- :waning_gibbous_moon: -->
<!-- :last_quarter_moon: -->
<!-- :waning_crescent_moon: -->
<!-- :last_quarter_moon_with_face: -->
<!-- :first_quarter_moon_with_face: -->
<!-- :moon: -->
