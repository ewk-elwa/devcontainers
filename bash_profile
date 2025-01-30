#!/bin/bash

source ~/.profile   # Load paths
source ~/.bashrc    # Load aliases

source ~/.git-prompt.sh   # Load the git branch prompt script

# Load bash-completion from brew
#[[ -r "/opt/homebrew/etc/profile.d/bash_completion.sh" ]] && . "/opt/homebrew/etc/profile.d/bash_completion.sh"

function parse_git_branch() {
     git branch 2> /dev/null | sed -e '/^[^*]/d' -e 's/* \(.*\)/(\1)/'
}

function _update_ps1() {
    PS1="$(powerline-shell $?)"
}

function _update_mux_ps1() {
  color_prompt=yes
  if [ "$color_prompt" = yes ]; then
      PS1="\[\033[0;31m\]\342\224\214\342\224\200\$([[ \$? != 0 ]] && echo \"[\[\033[0;31m\]\342\234\227\[\033[0;37m\]]\342\224\200\")[$(if [[ ${EUID} == 0 ]]; then echo '\[\033[01;31m\]root\[\033[01;33m\]@\[\033[01;96m\]\h'; else echo '\[\033[0;39m\]\u\[\033[01;33m\]@\[\033[01;96m\]\h'; fi)\[\033[0;31m\]]\342\224\200[\[\033[0;32m\]\w\[\033[0;31m\]] \[\e[91m\]\$(parse_git_branch)\[\e[00m\]\n\[\033[0;31m\]\342\224\224\342\224\200\342\224\200\342\225\274 \[\033[0m\]\[\e[01;33m\]\\$\[\e[0m\]"
  else
      PS1='┌──[\u@\h]─[\w]  \[\e[91m\]\$(parse_git_branch)\[\e[00m\]\n└──╼ \[\e[91m\]\$(parse_git_branch)\[\e[00m\]"\$ '
  fi

}

if [ "$TERM" != "linux" ]; then
  PROMPT_COMMAND="_update_mux_ps1; $PROMPT_COMMAND"
fi
