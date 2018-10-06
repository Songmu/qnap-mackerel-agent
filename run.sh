#!/bin/sh

DAEMON_MGR=/sbin/daemon_mgr
NAME=mackerel-agent
PROG=/usr/local/$NAME/bin/$NAME
CONF=/usr/local/$NAME/$NAME.conf

case "$1" in
    start)
        mypid=$(/bin/pidof $NAME)
        if [ ! -z $mypid ]; then
                exit 1
        else
            echo -n "Starting $NAME: "
            $DAEMON_MGR $NAME start "$PROG --conf=$CONF &amp;"
            exit 0
        fi
        ;;
    stop)
        echo -n "Stopping $NAME: "
        $DAEMON_MGR $NAME stop "$PROG"
        echo
        ;;
    *)
        echo "Usage: $1 start|stop"
        exit 1
        ;;
esac
