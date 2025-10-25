# Centrifugo gRPC golang sdk
[![Pipeline status](https://github.com/kachit/gocent-grpc/badges/master/pipeline.svg)](https://github.com/kachit/gocent-grpc/commits/master)
[![Coverage report](https://github.com/kachit/gocent-grpc/badges/master/coverage.svg)](https://github.com/kachit/gocent-grpc/commits/master)
[![Latest Release](https://github.com/kachit/gocent-grpc/-/badges/release.svg)](https://github.com/kachit/gocent-grpc/-/releases)

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