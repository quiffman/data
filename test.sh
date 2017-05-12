#!/bin/bash

set -x

errcount=0

error_handler () {
    echo "Trapped error - ${1:-"Unknown Error"}" 1>&2
    (( errcount++ ))       # or (( errcount += $? ))
}

trap error_handler ERR

go test ./CMT

exit $errcount

# vim: set ts=4 sw=4 tw=0 et:
