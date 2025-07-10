## Proto
proto.download: ## download latest proto file
	curl -o ./proto/api1.proto https://raw.githubusercontent.com/centrifugal/centrifugo/master/internal/apiproto/api.proto

proto.generate: ## generate proto files
	protoc --go_out=. --go-grpc_out=. --go-grpc_opt=paths=source_relative --go_opt=paths=source_relative proto/*.proto
## Tests
tests.unit: ## run unit tests
	go test ./...

tests.unit.coverage: ## run unit tests with coverage
	go test --coverprofile=coverage.out ./... ; go tool cover -func coverage.out ; go tool cover --html=coverage.out -o coverage.html

lint: ## run lint
	golangci-lint run
