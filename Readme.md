# TON SDK client Golang

## Preparations

One need to install
- [Rust](https://www.rust-lang.org/tools/install)
- [TON-SDK](https://github.com/tonlabs/TON-SDK) - *important* version 1.0.0-rc and compile it via `cargo build --release`

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

For examples see `cmd/cli/main.go` and run it via `go run ./cmd/cli`

## Tests

```shell script
export CGO_LDFLAGS="..."
task test
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

- https://eli.thegreenplace.net/2019/passing-callbacks-and-pointers-to-cgo/
- https://dev.to/mattn/call-go-function-from-c-function-1n3
- https://docs.ton.dev/86757ecb2/p/71d7a8-samples/t/35f373 - example how to use GraphQL