#!/usr/bin/tcsh

# XXX: this script is intended to be run from
# a fresh Digital Ocean droplet with FreeBSD
# Just run tcsh install_aphelion_bsd.sh

# upon its completion, you must either reset
# your terminal or run `source ~/.tcshrc`

# This assumes your installing it through tcsh as root.
# Change the relevant lines from tcsh to csh if your
# installing as a different user, along with changing the
# gopath.

# change this to a specific release or branch
set BRANCH=master
set REPO=github.com/libonomy/aphelion-staking

set GO_VERSION=1.12

sudo pkg update

sudo pkg install -y gmake
sudo pkg install -y git

# get and unpack golang
curl -O https://storage.googleapis.com/golang/go$GO_VERSION.freebsd-amd64.tar.gz
tar -xvf go$GO_VERSION.freebsd-amd64.tar.gz

# move go folder and add go binary to path
sudo mv go /usr/local
set path=($path /usr/local/go/bin)


# create the go directory, set GOPATH, and put it on PATH
mkdir go
echo "setenv GOPATH /root/go" >> ~/.tcshrc
setenv GOPATH /root/go
echo "set path=($path $GOPATH/bin)" >> ~/.tcshrc

source ~/.tcshrc

# get the code and move into repo
go get $REPO
cd "$GOPATH/src/$REPO"

# build & install master
git checkout $BRANCH
gmake get_tools
gmake install

# the binary is located in $GOPATH/bin
# run `source ~/.profile` or reset your terminal
# to persist the changes