# grpc
<!--
[![Go Report Card](https://goreportcard.com/badge/github.com/zcong1993/grpc)](https://goreportcard.com/report/github.com/zcong1993/grpc)
-->

> grpc-go example

## install

install [protobuf](https://github.com/google/protobuf) first

```sh
$ git clone https://github.com/zcong1993/grpc.git
$ dep ensure
# or
$ go get -v ./...
```

## run

```sh
# server
# go run service/main.go
# client
$ go run cmd/client.go
# http server
$ go run http/server.go
```

## update proto

modify `echo/echo.proto` then run `$ ./update.sh`

## License

MIT &copy; zcong1993
