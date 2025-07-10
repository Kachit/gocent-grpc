# Centrifugo gRPC golang sdk
[![Pipeline status](https://jgit.me/proxy/libraries/cf-rpc-go/badges/master/pipeline.svg)](https://jgit.me/proxy/libraries/cf-rpc-go/commits/master)
[![Coverage report](https://jgit.me/proxy/libraries/cf-rpc-go/badges/master/coverage.svg)](https://jgit.me/proxy/libraries/cf-rpc-go/commits/master)
[![Latest Release](https://jgit.me/proxy/libraries/cf-rpc-go/-/badges/release.svg)](https://jgit.me/proxy/libraries/cf-rpc-go/-/releases)

### Used by
- https://jgit.me/proxy/workers-synchronizer

## Install
Add `GOSUMDB=off` to envs

Add to `go.mod` this line
```
replace jgit.me/proxy/libraries/cf-rpc-go => jgit.me/proxy/libraries/cf-rpc-go.git v0.0.4
```
```shell
go mod tidy
```
```shell
go get -u jgit.me/proxy/libraries/cf-rpc-go
```

## Usage (Client)
```go
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "jgit.me/proxy/libraries/cf-rpc-go"
	"log"
)

func main() {
	conn, err := grpc.Dial(
		"127.0.0.1:10000",
		grpc.WithTransportCredentials(insecure.NewCredentials()), 
		grpc.WithPerRPCCredentials(pb.KeyAuth{Key: "qwerty"}),
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req := &pb.PresenceRequest{Channel: "worker:public"}
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