Meus arquivos ponto (famosos _dotfiles_)

## :underage: Instalação

Pré-requisito: [GNU Stow](https://www.gnu.org/software/stow/)


### :wrench: Personalizações

Arquivos `.zsh` em `$HOME/.zsh_extras` serão automaticamente _sourced_ quando o login for efetuado.


### :wrench: Tmux

* `~/.tmux.conf` configurações do tmux
* `~/.tmux.conf.user` crie esse arquivo para configurações pessoais
* `CTRL` + `a` como tecla padrão
* `~/bin/tmux_dev_env.sh` exemplo de automatização na abertura do tmux (necessita `~/bin` presente no `$PATH`)

```bash
stow tmux
```


### :nuts_and_bolt: Gitconfig

* `~/.gitconfig` várias configurações do git, e definição de arquivos padrão:
* `~/.gitconfig.user` crie esse arquivo para configurações pessoais
* `~/.gitconfig.commit-template.txt` como o nome indica
* `~/.gitignore` para ignores file

```bash
stow gitconfig
git config -l --show-origin  # mostra configurações e origem
```


### :wrench: Vim

* `~/.vimrc` configurações do vim
* `~/.vimrc.before` crie esse arquivo para configurações pessoais
* plugins recomendado: [NERDTree](https://github.com/preservim/nerdtree) (`CTRL` + `n`)

```bash
stow vim

# vim versão > 8
git clone https://github.com/preservim/nerdtree.git ~/.vim/pack/vendor/start/nerdtree
vim -u NONE -c "helptags ~/.vim/pack/vendor/start/nerdtree/doc" -c q
```


### :wrench: Alacritty

* `~/.alacritty` configurações padrão do terminal Alacritty

```bash
stow alacritty
```


### :hammer: ZSH

Pré-requisito: [Z Shell](http://zsh.sourceforge.net/Doc/Release/Introduction.html)

```bash
git clone git@github.com:crineu/arquivosponto.git ~/arquivosponto
cd ~/arquivosponto

stow zsh
chsh -s /usr/bin/zsh
```

Clona repositório, define o _shell_ como ZSH e habilita:

* Prompt com git status na linha de comando e alias:
    - `gba`, `gbb`: branchs e conexões remotas;
    - `gst`, `gsc`: git status, git commit
    - `glg`, `gld`, `glb`, `glc` - distintos tipos de log
* ZMV!
    - `zmv  'batata-(*)' 'banana.$1'` <-- :scream:


#### Ordem de inicialização dos arquivos:

* `.zshenv`: invocado sempre; deve ser conciso e meramente inicializar as variáveis necessárias;
* `.zlogin`: carregado em _login shells_ após `.zshrc`; compila _zcompdump_ em segundo plano pois é processo lento e feito 1x / login;
* `.zprofile` : similar a `.zlogin` mas carregado antes de `.zshrc` (`.zprofile` e  `.zshrc` são ignorados em _non-login non-interactive shells - so I learned a trick from Prezto that declares environment variables in _.zprofile_ and uses .zshenv_ to source _.zprofile_ (e.g. .zprofile and .zshenv). This way, non-login non-interactive shells will receive proper variable initialisations).
* `.zshrc`: carregado em _interactive shells_; contém a parte principal das configurações do ZSH.


### :wrench: ASDF

Recomenda-se a instalação do [ASDF](https://asdf-vm.com/#/core-manage-asdf-vm). Após:

```bash
stow asdf

asdf plugin-list-all
asdf plugin-add {ruby, rust, docker, tmux}
asdf install {ruby, rust, docker, tmux} _version_number
```

### :hammer: VS Code

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
