# Docker alias and help functions

alias dkls='docker image ls'
alias dkps='docker container ps -a'
alias dkvl='docker volume ls'

# running lazydocker portainer hml
alias portainer-hml='docker \
  --host tcp://150.162.1.72:2376 --tls \
  run --rm -it \
  --name sisacad-tui \
  -v /var/run/docker.sock:/var/run/docker.sock \
  lazyteam/lazydocker'
