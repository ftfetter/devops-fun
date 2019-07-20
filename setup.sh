#!/bin/bash

###############################################################
# Script to setup the GoLang Environment on Vagrant Ubuntu VM #
###############################################################

echo "Start provisioning..."

#saving provisioning date
date > /etc/vagrant_provisioned_at

#intalling git
sudo apt install -y git

#downloading golang 1.12.6
curl -O https://storage.googleapis.com/golang/go1.12.6.linux-amd64.tar.gz

#unpatching
sudo tar -xvf go1.12.6.linux-amd64.tar.gz

#moving to home directory
sudo mv go /usr/local/go

#removing source file
rm go1.12.6.linux-amd64.tar.gz

#configuring Go env
echo "export GOROOT=/usr/local/go" >> ~/.profile
export GOROOT=/usr/local/go
echo "export PATH=$PATH:$GOROOT/bin" >> ~/.profile
export PATH=$PATH:$GOROOT/bin
echo "export GOPATH=/home/vagrant/go" >> ~/.profile
export GOPATH=/home/vagrant/go

#creating service directory
mkdir -p /home/vagrant/go/src/calc-service