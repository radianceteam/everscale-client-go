# TON SDK client Golang

[![Awesome Badges](https://img.shields.io/badge/SDK_version-1.5.2-green.svg)](https://github.com/tonlabs/TON-SDK/tree/1.5.2)
[![Awesome Badges](https://img.shields.io/badge/TON_version-0.24.8-green.svg)](https://hub.docker.com/layers/tonlabs/local-node/0.24.8/images/sha256-62239cb2b215cbca7e8792812e27fa293727cfd8b17d3e58523c8a15a3673504?context=explore)
[![Chat Telegram](https://img.shields.io/badge/chat-Telegram-9cf.svg)](https://t.me/RADIANCE_TON_SDK)
![CI tests and linters](https://github.com/radianceteam/ton-client-go/workflows/CI/badge.svg)

## Preparations

One need to install
- [Rust](https://www.rust-lang.org/tools/install)
- [TON-SDK](https://github.com/tonlabs/TON-SDK) - *important* version `1.5.2` and compile it via `cargo build --release`

## Run

One need to specify compiled DLL directory path:
```shell script
export CGO_LDFLAGS="-L/path-to-installation/TON-SDK/target/release/deps/ -lton_client"
go build ./cmd/cli
go run ./cmd/cli
# or
task run
```

On Linux one need also provide search path for DLL loader:
```shell script
export LD_LIBRARY_PATH=/path-to-installation/TON-SDK/target/release/deps/
```
## Wrapper usage

For examples see `cmd/cli/main.go` and run it via `go run ./cmd/cli`.
This wrapper covers 100% of functionality for TON-SDK.
All non-generated code has test coverage at least of 70% - one can see it via `task coverage`.

## Tests

```shell script
export CGO_LDFLAGS="..."
docker run -d --name local-node -p80:80 tonlabs/local-node:0.24.8
task test # tests without node
task full_test # tests including with node
task coverage # full_test with coverage
```

## Development

See available task commands via `task` without arguments.
You need to install:
- [golangci-lint](https://github.com/golangci/golangci-lint).
- [Taskfile](https://taskfile.dev/) (optional)
To attach git hooks run `task attach_hooks`

### Code generation

Most of the code is generated via script in `./cmd/gen/gen.go` from spec `api-spec.json`
one can regenerate it via `task generate`.

## Useful reading

- https://github.com/tonlabs/TON-SDK/blob/1.5.2/docs/app_objects.md
- https://eli.thegreenplace.net/2019/passing-callbacks-and-pointers-to-cgo/
- https://dev.to/mattn/call-go-function-from-c-function-1n3
- https://docs.ton.dev/86757ecb2/p/71d7a8-samples/t/35f373 - example how to use GraphQL

## Methods and types with manual implementation
See more at `spec/types.go`
- processing.send_message
- processing.wait_for_transaction
- processing.process_message
- net.subscribe_collection
