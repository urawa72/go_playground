#!/bin/bash

if ! which vim >dev/null 2>&1; then
  cd ~
  git clone https://github.com/urawa72/dotfiles.git
  cd dotfiles
  ./setup_linux.sh
fi
exec /bin/bash
