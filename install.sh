#!/bin/sh
set -e

install_path=${MACKEREL_INSTALL_PATH:-/share/MD0_DATA/.mackerel-agent}
install -d "$install_path"
cp -rf "$(dirname "$0")/" "$install_path"

# curl -o /etc/ssl/certs/ca-certificates.crt https://curl.haxx.se/ca/cacert.pem

cd $install_path
if [ ! -f ./mackerel-agent.conf ]; then
  bin/mkrinst agentconf "$install_path"
  install -m 0600 mackerel-agent.conf.sample mackerel-agent.conf
  while [ "$MACKEREL_APIKEY" = "" ]; do
    read -rp "input Mackerel APIKEY: " MACKEREL_APIKEY
  done
  bin/mackerel-agent init "-apikey=$MACKEREL_APIKEY" -conf ./mackerel-agent.conf
fi

bin/mkrinst qpkgconf "$install_path"
