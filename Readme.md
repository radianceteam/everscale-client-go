# TON SDK client Golang

## Preparations

One need to install
- [Golang](https://golang.org/doc/install)
- [Taskfile](https://taskfile.dev/)
- [TON-SDK](https://github.com/tonlabs/TON-SDK) 

## Run 

One need to specify compiled DLL directory path:
```shell script
export CGO_LDFLAGS="-L/Users/risentveber/Projects/TON-SDK/target/release/deps/ -lton_client"
go build ./cmd/cli
go run ./cmd/cli
# or
task run
```
## Tests

```shell script
export CGO_LDFLAGS="..."
task test
```

## Development

See available task commands via `task` without arguments.
You need to install [golangci-lint](https://github.com/golangci/golangci-lint).
To attach git hooks run `task attach_hooks`

## Useful reading

- https://eli.thegreenplace.net/2019/passing-callbacks-and-pointers-to-cgo/
- https://dev.to/mattn/call-go-function-from-c-function-1n3