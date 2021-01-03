#!/bin/sh

rm -f /tmp/libton_client.so.gz /tmp/libton_client.so
wget http://sdkbinaries-ws.tonlabs.io/tonclient_1_linux.gz -O /tmp/libton_client.so.gz
gzip -d /tmp/libton_client.so.gz

