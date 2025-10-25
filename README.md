# Centrifugo gRPC golang sdk
[![Go Test](https://github.com/Kachit/gocent-grpc/actions/workflows/tests.yml/badge.svg)](https://github.com/Kachit/gocent-grpc/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/kachit/gocent-grpc)](https://goreportcard.com/report/github.com/kachit/gocent-grpc)
[![Go Version](https://img.shields.io/github/go-mod/go-version/Kachit/gocent-grpc)](https://go.dev/doc/go1.24)
[![Release](https://img.shields.io/github/v/release/Kachit/gocent-grpc.svg)](https://github.com/Kachit/gocent-grpc/releases)
[![GoDoc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kachit/gocent-grpc)
[![License](https://img.shields.io/github/license/Kachit/gocent-grpc)](https://github.com/Kachit/gocent-grpc/blob/master/LICENSE)

## Install
```shell
go get -u github.com/kachit/gocent-grpc
```

## Usage (Client)

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/kachit/gocent-grpc/auth"
	pb "github.com/kachit/gocent-grpc/pkg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(
		"127.0.0.1:10000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(auth.KeyAuth{Key: "qwerty"}),
	)
	if err != nil {
		log.Fatalln(err)
	}
	
	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req := &pb.PresenceRequest{Channel: "channel-name"}
	//
	client := pb.NewCentrifugoApiClient(conn)
	
	result, err := client.Presence(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(result)
}
```

## Proto
https://raw.githubusercontent.com/centrifugal/centrifugo/master/internal/apiproto/api.proto

## ProtoGen Golang
```shell
protoc --go_out=. --go-grpc_out=. --go-grpc_opt=paths=source_relative --go_opt=paths=source_relative proto/*.proto
```

## Linters ##
```bash
golangci-lint run
```