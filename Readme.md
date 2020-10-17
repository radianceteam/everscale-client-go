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

## Useful reading

- https://eli.thegreenplace.net/2019/passing-callbacks-and-pointers-to-cgo/
- https://dev.to/mattn/call-go-function-from-c-function-1n3