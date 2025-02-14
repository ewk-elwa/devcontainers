#!/bin/sh
echo "install tools needed for tests"
apk add --no-cache docker-cli bash curl openssl git
curl -s https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash
curl -sL https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 | bash
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
chmod +x kubectl
mv kubectl /usr/local/bin/

