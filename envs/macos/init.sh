#!/bin/bash

cd ~/

# Goバージョン設定
GO_VERSION=

# Goのインストール（既存のバージョンのGoの削除）
sudo rm -rf /usr/local/go && tar -C /usr/local -xzf go$GO_VERSION.linux-amd64.tar.gz

# .tarファイルの消去
sudo rm -rf go$GO_VERSION.linux-amd64.tar.gz

# PATHの設定
echo 'export PATH=$PATH:/usr/local/go/bin' >> .bashrc

# GOPATHの設定
echo 'export GOPATH=$HOME/.go' >> .bashrc
echo 'export PATH=$PATH:$GOPATH/bin' >> .bashrc

# golangci-lintのインストール
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.46.2

