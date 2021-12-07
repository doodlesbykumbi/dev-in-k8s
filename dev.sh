#!/usr/bin/env bash

/go/bin/reflex \
    -r "\.go$" \
    -R "_test\.go$" \
    -s -- \
        /go/bin/dlv debug \
            --listen=:40000 \
            --headless=true \
            --api-version=2 \
            --continue \
            --accept-multiclient
