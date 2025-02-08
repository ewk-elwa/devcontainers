#!/bin/bash
export ARCH=amd64
wget -q https://get.helm.sh/helm-v3.7.2-linux-${ARCH}.tar.gz && \
    tar -zxvf helm-v3.7.2-linux-${ARCH}.tar.gz && \
    mv linux-${ARCH}/helm /usr/local/bin/helm && \
    rm -rf linux-${ARCH} helm-v3.7.2-linux-${ARCH}.tar.gz

wget -O /usr/local/bin/kubectl https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/${ARCH}/kubectl && \
    chmod +x /usr/local/bin/kubectl
wget -q -O - https://raw.githubusercontent.com/rancher/k3d/main/install.sh | bash

# install brew too
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"


# install helmify
brew install arttor/tap/helmify
