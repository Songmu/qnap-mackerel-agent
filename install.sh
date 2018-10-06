#!/bin/sh
set -e

install_path=/share/MD0_DATA/.mackerel-agent
install -d "$install_path"
cp -r "$(dirname "$0")/" "$install_path"
ln -s $install_path /usr/local/mackerel-agent

cd $install_path
if [ ! -f ./mackerel-agent.conf ]; then
  install mackerel-agent.conf.sample mackerel-agent.conf
  while [ "$MACKEREL_APIKEY" = "" ]; do
    read -rp "input Mackerel APIKEY: " MACKEREL_APIKEY
  done
  bin/mackerel-agent init "$MACKEREL_APIKEY" -conf ./mackerel-agent.conf
fi

qpkgconf=/etc/config/qpkg.conf
if ! grep '\[mackerel_agent\]' $qpkgconf > /dev/null 2>&1; then
    cat ./qpkg.conf.txt >> $qpkgconf
fi