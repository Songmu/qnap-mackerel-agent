#!/bin/sh
set -e

install_path=${MACKEREL_INSTALL_PATH:-/share/MD0_DATA/.mackerel-agent}
install -d "$install_path"
cp -rf "$(dirname "$0")/" "$install_path"

# curl -o /etc/ssl/certs/ca-certificates.crt https://curl.haxx.se/ca/cacert.pem

cd $install_path
if [ ! -f ./mackerel-agent.conf ]; then
  while [ "$MACKEREL_APIKEY" = "" ]; do
    read -rp "input Mackerel APIKEY: " MACKEREL_APIKEY
  done
  bin/mkrinst agentconf "$install_path" "$MACKEREL_APIKEY"
fi

bin/mkrinst qpkgconf "$install_path"
