#!/usr/bin/env sh

[ "$PINGME_USE_SERVICE" ]
: ${PINGME_USE_SERVICE:=""}

pingme $PINGME_USE_SERVICE