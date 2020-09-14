FROM golang:latest

ENV TERM=xterm-256color

RUN apt-get update && \
    apt-get install -y \
    curl \
    git \
    zsh \
    powerline \
    fonts-powerline

# for vim build
RUN apt-get install -y \
  git \
  gettext \
  libtinfo-dev \
  libacl1-dev \
  libgpm-dev \
  build-essential \
  libperl-dev \
  python-dev \
  python3-dev \
  ruby-dev \
  lua5.2 \
  liblua5.2-dev \
  autoconf \
  automake \
  cproto \
  libxmu-dev \
  libgtk-3-dev \
  libxpm-dev

RUN cd /tmp && \
  git clone https://github.com/vim/vim.git && \
  cd vim && \
  ./configure \
    --with-features=huge \
    --enable-multibyte \
    --enable-cscope \
    --enable-fontset \
    --enable-fail-if-missing \
    --enable-pythoninterp=dynamic \
    --enable-python3interp=dynamic \
    --enable-rubyinterp=dynamic \
    --enable-gui=true \
    --enable-gui=gtk3 \
    --with-x && \
  make && make install

CMD ["/usr/bin/zsh"]
