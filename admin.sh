#!/bin/bash

SERVER="apiserver"
BASE_DIR=$PWD
INTERVAL=2

ARGS=""

function start()
{
    if [ "`pgrep $SERVER -u $UID`" != "" ]; then
        echo "$SERVER aleady running"
        exit 1
    fi
    nohup $BASE_DIR/$SERVER $ARGS server &>/dev/nill &
    echo "sleeping..." && sleep $INTERVAL

    if [ "`pgrep $SERVER -u $UID`" != "" ]; then
        echo "$SERVER is running"
    else
        echo "$SERVER is not running"
    fi
}

function status()
{
    if [ "`pgrep $SERVER -u $UID`" != "" ];then
        echo "$SERVER is running"
    else
        echo "$SERVER is not running"
    fi
}


function stop()
{
    if [ "`pgrep $SERVER -u $UID`" != "" ];then
        kill -9 `pgrep $SERVER -u $UID`
    fi
    echo "sleeping..." && sleep $INTERVAL
    if [ "`pgrep $SERVER -u $UID`" != "" ]; then
        echo "$SERVER stop failed"
        exit 1
    fi
}



case "$1" in
    'start')
    start
    ;;
    'stop')
    stop
    ;;
    'status')
    status
    ;;
    'restart')
    stop && start
    ;;
    *)
    echo "usage:$0 {start|stop|restart|status}"
    exit 1
    ;;
esac