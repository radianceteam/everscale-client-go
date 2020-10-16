# TON SDK client Golang

## Run 

One need to specify compiled DLL directory path:
```
export CGO_LDFLAGS="-L/Users/risentveber/Projects/TON-SDK/target/release/deps/ -lton_client"
go build ./cmd/cli
go run ./cmd/cli
```
## Tests

## Useful reading

- https://eli.thegreenplace.net/2019/passing-callbacks-and-pointers-to-cgo/
- https://dev.to/mattn/call-go-function-from-c-function-1n3