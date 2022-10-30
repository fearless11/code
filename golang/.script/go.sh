#!/bin/bash
# date: 2022/10/30
# auth: fearless11
# desc: install go
# doc: https://gomirrors.org/

version=1.18.2
# linux
package=go${version}.linux-amd64.tar.gz
# mac
# package=go${version}.darwin-amd64.tar.gz
url="https://gomirrors.org/dl/go/${package}"
dir="/usr/local/go"
githubproxy="https://ghproxy.com/"

install() {
    mkdir -p ${dir} ${dir}/gopath &>/dev/null
    cd ${dir} && wget ${url}
    mkdir -p go${version} >/dev/null && tar xf ${package} -C go${version}
    ln -s ${dir}/go${version}/go ${dir}/goroot
    cp /etc/profile /home/profile-$(date +%F-%T)

    cat <<EOF >>/etc/profile
export GOROOT=/usr/local/go/goroot
export GOPATH=/usr/local/go/gopath
export GOPROXY=https://proxy.golang.com.cn,direct
export GOPRIVATE="*.company.com"
export GO111MODULE=on
export PATH=\$GOROOT/bin:\$GOPATH/bin:\$PATH
EOF
    source /etc/profile

    # install go tools
    # set -mod fix missing go.sum entry 
    go env -w "GOFLAGS"="-mod=mod"
    mkdir -p $GOPATH/src/golang.org/x
    cd $GOPATH/src/golang.org/x
    git clone ${githubproxy}https://github.com/golang/tools.git tools
    git clone ${githubproxy}https://github.com/golang/lint.git lint
    git clone ${githubproxy}https://github.com/golang/net.git net
    cd $GOPATH/src/golang.org/x/tools/cmd && go install ...
    cd $GOPATH/src/golang.org/x/lint && go install ...
    cd $GOPATH/src/golang.org/x/net && go install ...
    go env -w "GOFLAGS"=""

    cd ${dir}
    go version
    echo "please exec: source /etc/profile"
}

uninstall() {
    sed -i '/GOROOT/d' /etc/profile
    sed -i '/GOPATH/d' /etc/profile
    sed -i '/GOPROXY/d' /etc/profile
    sed -i '/GOPRIVATE/d' /etc/profile
    sed -i '/GO111MODULE/d' /etc/profile
    rm -rf ${dir}
}

update() {
    cd ${dir}
    wget $url
    mkdir -p go${version} >/dev/null && tar xf ${package} -C go${version}
    unlink ${dir}/goroot && ln -s ${dir}/go${version}/go ${dir}/goroot
}


case $1 in
install)
    install
    ;;
uninstall)
    uninstall
    ;;
update)
    update
    ;;
*)
    echo "$0 install|uninstall|update"
    ;;
esac
