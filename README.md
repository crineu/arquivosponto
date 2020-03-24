Meus arquivos ponto (famosos _dotfiles_)

## :wrench: Instalação

Pré-requisito: [GNU Stow](https://www.gnu.org/software/stow/)

```bash
cd ~
git clone git@github.com:crineu/arquivosponto.git
cd arquivosponto

stow --verbose zsh
chsh -s zsh
```

Clona repositório, define o _shell_ como ZSH.


### :frog: Gitconfig

Configurações do git, e definição de arquivos padrão:

* `~/.gitconfig.user` configurações pessoais
* `~/.gitconfig.commit-template.txt` como o nome indica
* `~/.gitignore` para ignores file


```
stow gitconfig
```


### :monkey: Tmux

Configurações do tmux, e definição de arquivos padrão:

* `~/.tmux.conf.user` configurações pessoais


```
stow tmux
```

