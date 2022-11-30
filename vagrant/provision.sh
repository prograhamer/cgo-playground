#!/bin/bash

set -euf -o pipefail

apt-get update
apt-get -y upgrade

# Install build essential and debugger
apt-get install -y build-essential gdb

# Install Go 1.19
if [[ ! -d /usr/local/go ]]; then
   wget https://go.dev/dl/go1.19.3.linux-amd64.tar.gz
   tar -xzf go1.19.3.linux-amd64.tar.gz
   sudo mv go /usr/local

   sudo echo 'export PATH=$PATH:/usr/local/go/bin' >> /etc/profile
fi
