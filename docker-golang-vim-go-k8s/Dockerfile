#FROM golang:1.9.2
FROM golang:1.11

# Install vim
RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y vim && \
    apt-get clean

RUN DEBIAN_FRONTEND=noninteractive apt-get install -y vim-nox && \
    apt-get clean


# VIM plugin for go
RUN git clone https://github.com/VundleVim/Vundle.vim.git ~/.vim/bundle/Vundle.vim


COPY kubectl /usr/bin/kubectl
COPY vimrc /root/.vimrc
RUN cd /go
RUN echo -ne '\n' |vim +PluginInstall +qall
RUN echo -ne '\n' |vim  +GoInstallBinaries +qall; exit 0
RUN go get -u github.com/kardianos/govendor


# Glide for go dependency management
#RUN curl https://glide.sh/get | sh
#RUN git clone https://github.com/fatih/vim-go.git ~/.vim/bundle/vim-go
