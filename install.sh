#!/bin/sh
set -e

install_path=/share/MD0_DATA/.mackerel-agent
install -d $install_path
cp -r $(dirname $0) $install_path
ln -s $install_path /usr/local/mackerel-agent

cd $install_path
while [ "$MACKEREL_APIKEY" = "" ]; do
  read -p "input Mackerel APIKEY: " MACKEREL_APIKEY
done
bin/mackerel-agent init $MACKEREL_APIKEY -conf ./mackerel-agent.conf

qpkgconf=/etc/config/qpkg.conf
if ! grep '\[mackerel_agent\]' $qpkgconf 2>&1 > /dev/null; then
    cat ./qpkg.conf.txt >> $qpkgconf
fi
