#!/bin/sh

DAEMON_MGR=/sbin/daemon_mgr
NAME=mackerel-agent
PROG_DIR=$(cd "$(dirname "$0")"; pwd)
PROG="$PROG_DIR/bin/$NAME"
CONF="$PROG_DIR/$NAME.conf"

start() {
  mypid=$(/bin/pidof $NAME)
  if [ ! -z "$mypid" ]; then
    exit 1
  fi

  /bin/echo -n $"Starting $NAME: "
  $DAEMON_MGR $NAME start "$PROG --conf=$CONF 2>&1 | logger -t $NAME &"
  /bin/echo "OK"
}

stop() {
  /bin/echo -n $"Stopping $NAME: "
  $DAEMON_MGR $NAME stop "$PROG"
  /bin/echo "OK"
}

restart() {
  stop
  /bin/sleep 1
  start
}

case "$1" in
  start)
    start
    ;;
  stop)
    stop
    ;;
  restart)
    restart
    ;;
  *)
    /bin/echo $"Usage: $0 {start|stop|restart}"
    exit 1
    ;;
esac
