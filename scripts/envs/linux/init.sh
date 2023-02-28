#!/bin/bash

cd ~/


# Goバージョン設定
NEW_GO_VERSION="1.20"

# 既存のGoの削除
sudo apt-get purge golang*

# Goのインストール
wget https://dl.google.com/go/go$NEW_GO_VERSION.linux-amd64.tar.gz

# 権限の変更
sha256sum go$NEW_GO_VERSION.linux-amd64.tar.gz

# tarファイルの展開
sudo rm -rf /usr/local/go && tar -C /usr/local -xzf go$NEW_GO_VERSION.linux-amd64.tar.gz

# .tarファイルの消去
sudo rm -rf go$GO_VERSION.linux-amd64.tar.gz

# PATHの設定
echo 'export PATH=$PATH:/usr/local/go/bin' >> .bashrc

# GOPATHの設定
echo 'export GOPATH=$HOME/.go' >> .bashrc
echo 'export PATH=$PATH:$GOPATH/bin' >> .bashrc

# golangci-lintのインストール
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.46.2
