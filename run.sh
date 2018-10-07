#!/bin/sh

DAEMON_MGR=/sbin/daemon_mgr
NAME=mackerel-agent
PROG=/share/MD0_DATA/.$NAME/bin/$NAME
CONF=/share/MD0_DATA/.$NAME/$NAME.conf

case "$1" in
  start)
    mypid=$(/bin/pidof $NAME)
    if [ ! -z "$mypid" ]; then
      exit 1
    else
      echo "Starting $NAME: "
      echo
      $DAEMON_MGR $NAME start "$PROG --conf=$CONF &"
      exit 0
    fi
    ;;
  stop)
    echo "Stopping $NAME: "
    echo
    $DAEMON_MGR $NAME stop "$PROG"
    echo
    ;;
  *)
    echo "Usage: $1 start|stop"
    exit 1
    ;;
esac
