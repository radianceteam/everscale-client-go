# TON SDK client Golang

[![EVER-SDK](https://img.shields.io/badge/EVER_SDK-1.41.0-green.svg)](https://github.com/tonlabs/EVER-SDK/tree/1.41.0)
[![TON local-node docker image](https://img.shields.io/badge/TON_local_node-0.33.0-green.svg)](https://hub.docker.com/layers/local-node/tonlabs/local-node/0.33.0/images/sha256-40a544432089f2d730f3dd0542fcb59bef6c5192b5d374fdae3a7ea6417070c4)
[![Chat Telegram](https://img.shields.io/badge/chat-Telegram-9cf.svg)](https://t.me/RADIANCE_EVER_SDK)
[![Documentation](https://godoc.org/github.com/radianceteam/everscale-client-go/client?status.svg)](https://godoc.org/github.com/radianceteam/everscale-client-go/client)
![CI tests and linters](https://github.com/radianceteam/everscale-client-go/workflows/CI/badge.svg)

## Preparations

One needs to install
- [Golang](https://golang.org/doc/install)
- [EVER-SDK](https://github.com/tonlabs/EVER-SDK#download-precompiled-binaries) - download precompiled binaries and extract them

### EVER-SDK installation - Mac OS
```bash
export EVER_SDK_INSTALLATION_PATH=`pwd`/tmp # example - should be specified as absolute path
# clean previous installation in case of reinstalling
rm -f $EVER_SDK_INSTALLATION_PATH/libton_client.dylib.gz $EVER_SDK_INSTALLATION_PATH/libton_client.dylib
# download binaries
wget https://binaries.tonlabs.io/tonclient_1_darwin.gz -O $EVER_SDK_INSTALLATION_PATH/libton_client.dylib.gz
# extract binaries
gzip -d $EVER_SDK_INSTALLATION_PATH/libton_client.dylib.gz
# make extracted file executable
chmod +x $EVER_SDK_INSTALLATION_PATH/libton_client.dylib
# set loading by absolute path
install_name_tool -id $EVER_SDK_INSTALLATION_PATH/libton_client.dylib $EVER_SDK_INSTALLATION_PATH/libton_client.dylib

# Better to add this to ~/.bashrc or ~/.zshrc to DRY in terminal each time you use it
export CGO_LDFLAGS="-L$EVER_SDK_INSTALLATION_PATH -lton_client"
```

### EVER-SDK installation - Linux
```bash
EVER_SDK_INSTALLATION_PATH=`pwd`/tmp # example - should be specified as absolute path
# clean previous installation in case of reinstalling
rm -f $EVER_SDK_INSTALLATION_PATH/libton_client.so.gz $EVER_SDK_INSTALLATION_PATH/libton_client.so
# download binaries
wget https://binaries.tonlabs.io/tonclient_1_linux.gz -O $EVER_SDK_INSTALLATION_PATH/libton_client.so.gz
# extract binaries
gzip -d $EVER_SDK_INSTALLATION_PATH/libton_client.so.gz
# make extracted file executable
chmod +x $EVER_SDK_INSTALLATION_PATH/libton_client.so

# Better to add this to ~/.bashrc or ~/.zshrc to DRY in terminal each time you use it
export LD_LIBRARY_PATH=$EVER_SDK_INSTALLATION_PATH:$LD_LIBRARY_PATH
export CGO_LDFLAGS="-L$EVER_SDK_INSTALLATION_PATH -lton_client"
```

## Run

One needs to specify compiled DLL directory path:
```shell script
export CGO_LDFLAGS="-L$EVER_SDK_INSTALLATION_PATH -lton_client"
go build ./cmd/cli
go run ./cmd/cli
# or
task run
```

On Linux one needs to provide search path for DLL loader:
```shell script
export LD_LIBRARY_PATH=$EVER_SDK_INSTALLATION_PATH:$LD_LIBRARY_PATH
```

## Wrapper usage

For examples see `cmd/cli/main.go` and run it via `go run ./cmd/cli`.
This wrapper covers 100% of functionality for EVER-SDK.
All non-generated code has test coverage at least of 70% - one can see it via `task coverage`.

## Tests

```shell script
export CGO_LDFLAGS="-L$EVER_SDK_INSTALLATION_PATH -lton_client"
docker run -d --name local-node -p80:80 tonlabs/local-node:0.33.0
task test # tests without node
task full_test # tests including with node
task coverage # full_test with coverage
```

## Development

You need to install:
- [golangci-lint](https://github.com/golangci/golangci-lint).
- [Taskfile](https://taskfile.dev/) (optional)

See available task commands via `task` without arguments.
To attach git hooks run `task attach_hooks`

### Code generation

Most of the code is generated via script in `./cmd/gen/gen.go` from spec `api-spec.json`
one can regenerate it via `task generate`.

## Useful reading

- https://medium.com/@donblas/fun-with-rpath-otool-and-install-name-tool-e3e41ae86172 - about DLL linking
- https://github.com/tonlabs/EVER-SDK/blob/1.5.2/docs/app_objects.md
- https://eli.thegreenplace.net/2019/passing-callbacks-and-pointers-to-cgo/
- https://dev.to/mattn/call-go-function-from-c-function-1n3
- https://docs.ton.dev/86757ecb2/p/71d7a8-samples/t/35f373 - example how to use GraphQL

## Methods and types with manual implementation

See more at `spec/types.go`
- processing.send_message
- processing.wait_for_transaction
- processing.process_message
- net.subscribe_collection
