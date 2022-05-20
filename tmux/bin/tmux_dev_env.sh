#!/bin/zsh

SESSIONNAME="dev"
tmux has-session -t $SESSIONNAME &> /dev/null

if [ $? != 0 ] 
 then
    cd ~/setic/
    tmux new-session -s $SESSIONNAME -n docker -d
    tmux new-window -t $SESSIONNAME -n 'git' -d
    tmux new-window -t $SESSIONNAME -n 'api'      -c '/home/crineu/setic/proj1/api' -d 
    tmux new-window -t $SESSIONNAME -n 'frontend' -c '/home/crineu/setic/proj1/frontend' -d
    tmux new-window -t $SESSIONNAME -n 'dkls' -c '/home/crineu' -d
    tmux new-window -t $SESSIONNAME -n 'wiki' -c '/home/crineu/setic/proj1.wiki' -d
fi

tmux attach -t $SESSIONNAME
