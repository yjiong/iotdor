#!/bin/bash
#########################################################################
# File Name: entryDelay.sh
# Author: yjiong
# mail: 4418229@qq.com
# Created Time: 2021-01-15 22:57:08
#########################################################################
: ${SLEEP_SECOND:=1}

wait_for() {
    echo Waiting for $1 to listen on $2...
    while ! nc -z $1 $2; do echo waiting...$1:$2; sleep $SLEEP_SECOND; done
}

declare DEPENDS
declare CMD

while getopts "d:c:" arg
do
    case $arg in
        d)
            DEPENDS=$OPTARG
            ;;
        c)
            CMD=$OPTARG
            ;;
        *)
            shift 1
            ;;
    esac
done

for var in ${DEPENDS//,/ }
do
    host=${var%:*}
    port=${var#*:}
    wait_for $host $port
done

eval $CMD

