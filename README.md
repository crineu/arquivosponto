Meus arquivos ponto (famosos _dotfiles_)

## :underage: Instalação

Pré-requisitos: [GNU Stow](https://www.gnu.org/software/stow/) e [Z Shell](http://zsh.sourceforge.net/Doc/Release/Introduction.html)

```bash
cd ~
git clone git@github.com:crineu/arquivosponto.git
cd arquivosponto

stow --verbose zsh
chsh -s zsh
```

Clona repositório, define o _shell_ como ZSH e habilita:

* Prompt deveras bacana
* Git status na linha de comando
* Git alias
    - gba - Mostra branchs e conexões remotas;
    - gbb - Outro tipo de visualização de branches;
    - gst - git status
    - gc  - git commit
    - glg, gld, glb, glc - distintos tipos de log
* ZMV!
    - zmv  'dot-(*)' '.$1'


### :earth_asia: Personalizações

Qualquer arquivo `.zsh` colocado na pasta `$HOME/.zsh_extras` será automaticamente _sourced_ quando o login for efetuado.

Exemplo de arquivo para adicionar _aliases_ a comandos docker:

```bash
arquivo '$HOME/.zsh_extras/docker.zsh'

alias dkls='docker image ls'
```


### :earth_americas: Gitconfig

Várias configurações do git, e definição de arquivos padrão:

* `~/.gitconfig.user` crie esse arquivo para configurações pessoais 
* `~/.gitconfig.commit-template.txt` como o nome indica
* `~/.gitignore` para ignores file


```bash
stow gitconfig
```


### :earth_africa: Tmux

Configurações do tmux, e definição de arquivos padrão:

* `~/.tmux.conf.user` crie esse arquivo para configurações pessoais
* `CTRL` + `a` como tecla padrão

```bash
stow tmux
```


### :earth_americas: ASDF

Recomenda-se fortemente a instalação do [ASDF](https://asdf-vm.com/#/core-manage-asdf-vm).

```bash
stow asdf

asdf plugin-list-all
asdf plugin-add {ruby, rust, docker, tmux}
asdf install {ruby, rust, docker, tmux} _version_number
```

---

## :volcano: Remoção completa :volcano:

```bash
cd ~/arquivosponto

stow -D zsh
stow -D gitconfig
stow -D tmux
stow -D asdf

cd ..
rm -rf arquivosponto
```

<!-- :mushroom: -->
<!-- :gift: -->
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
