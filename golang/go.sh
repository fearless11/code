#!/bin/bash
# date: 2022-05-25
# auth: fearless11
# desc: install go

version=1.18.2
package=go${version}.linux-amd64.tar.gz
# package=go${version}.darwin-amd64.tar.gz
url=https://gomirrors.org/dl/go/${package}
soft=/tmp/${package}

install() {
    wget $url -O ${soft}
    mkdir -p /usr/local/go/go${version} >/dev/null
    tar xf ${soft} -C /usr/local/go${version}
    unlink /usr/local/go/goroot
    ln -s /usr/local/go/go${version}/go /usr/local/go/goroot
    mkdir -p /usr/local/go/gopath
    cp /etc/profile /home

    cat <<EOF >>/etc/profile
export GOROOT=/usr/local/go/goroot
export GOPATH=/usr/local/go/gopath
export GOPROXY=https://proxy.golang.com.cn,direct
export GOPRIVATE="*.code.oa.com,*.woa.com"
export GO111MODULE=on
export PATH=\$GOROOT/bin:\$GOPATH/bin:\$PATH
EOF

    source /etc/profile
    # go env
}

update() {
    wget $url -O ${soft}
    mkdir -p /usr/local/go/go${version} >/dev/null
    tar xf ${soft} -C /usr/local/go/go${version}
    unlink /usr/local/go/goroot
    ln -s /usr/local/go/go${version}/go /usr/local/go/goroot
}

set_goenv(){
    # set goproxy 
    go env -w GOPROXY=https://goproxy.io,direct
    go env -w GO111MODULE="on"
}

case $1 in
install)
    install
    ;;
update)
    update
    ;;
esac
