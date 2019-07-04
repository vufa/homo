FROM debian:stretch

WORKDIR /home/homo

COPY . /home/homo/homo/

# Golang env
ENV GOLANG_VERSION 1.12.6
ENV GOLANG_TAR_BALL go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/$GOLANG_TAR_BALL
ENV GOLANG_DOWNLOAD_SHA256 dbcf71a3c1ea53b8d54ef1b48c85a39a6c9a935d01fc8291ff2b92028e59913c
ENV GOPATH /home/homo/go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

# Install system dependence
RUN \
    apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates git wget tar sudo && \
    apt-get install -y --no-install-recommends gcc automake autoconf libtool build-essential && \
    apt-get install -y --no-install-recommends bison swig python-dev libpulse-dev portaudio19-dev libwebkit2gtk-4.0-dev

# Add user homo to sudo
RUN useradd -m homo && echo "homo:homo" | chpasswd && adduser homo sudo

# Install PocketSphinx
RUN \
    cd homo && \
    make deps

# Install Golang
RUN wget $GOLANG_DOWNLOAD_URL && \
    echo "$GOLANG_DOWNLOAD_SHA256  $GOLANG_TAR_BALL" | sha256sum -c - && \
    sudo tar -C /usr/local -xzf $GOLANG_TAR_BALL && \
    rm $GOLANG_TAR_BALL

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

RUN go env

# Build homo webview
RUN \
    cd homo && \
    make gen && \
    make webview

# Install python
RUN \
    wget https://www.python.org/ftp/python/3.6.8/Python-3.6.8.tgz && \
    tar xvf Python-3.6.8.tgz && \
    cd Python-3.6.8 && \
    ./configure --enable-optimizations --with-ensurepip=install && \
    make -j 8 && \
    sudo make altinstall

RUN python3.6 -V

# Install python dependencies
RUN \
    pip install virtualenv && \
    cd homo/nlu && \
    virtualenv --python=python3.6 env3.6 && \
    source env3.6/bin/activate && \
    pip install -r requirements.txt

# X11
#RUN export uid=1000 gid=1000 && \
#    mkdir -p /home/homo && \
#    echo "homo:x:${uid}:${gid}:homo,,,:/home/homo:/bin/bash" >> /etc/passwd && \
#    echo "homo:x:${uid}:" >> /etc/group && \
#    echo "homo ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/homo && \
#    chmod 0440 /etc/sudoers.d/homo && \
#    chown ${uid}:${gid} -R /home/homo

USER homo
ENV HOME /home/homo

VOLUME ["/home/homo/homo/conf", "/home/homo/homo/sphinx/en-us", "/home/homo/homo/sphinx/cmusphinx-zh-cn-5.2", "/home/homo/homo/nlu/models"]

CMD ["/home/homo/homo/nlu.sh", "/home/homo/homo/homo-webview"]
