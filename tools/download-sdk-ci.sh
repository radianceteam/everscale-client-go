#!/bin/sh

rm -f /tmp/libton_client.so.gz /tmp/libton_client.so
wget https://binaries.tonlabs.io/tonclient_1_linux.gz -O /tmp/libton_client.so.gz
gzip -d /tmp/libton_client.so.gz

