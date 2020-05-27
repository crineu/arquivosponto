# Docker alias and help functions

alias dkls='docker image ls'
alias dkps='docker container ps -a'
alias dkvl='docker volume ls'

alias ttt='echo ~'

# running lazydocker
alias lazydocker='docker \
  run --rm -it \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v /~/.config/lazydocker:/.config/jesseduffield/lazydocker \
  lazyteam/lazydocker'

# running lazydocker portainer hml
alias portainer-hml='docker \
  --host tcp://150.162.1.72:2376 --tls \
  run --rm -it \
  --name sisacad-tui \
  -v /var/run/docker.sock:/var/run/docker.sock \
  lazyteam/lazydocker'
