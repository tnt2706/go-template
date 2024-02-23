.PHONY:	start gql-gen gen-grpc kill-port

start:
	air /cmd/main.go

gen-cal:
    protoc \
    --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pkg/calculator/proto/calculator.proto

gen-grpc:
    protoc --go_out=. --go_opt=Mprotos/calculator.proto=pb \
    --go-grpc_out=. --go-grpc_opt=Mprotos/calculator.proto=pb \
    pkg/calculator/proto/calculator.proto

gql-gen:
    # generate tools.go
	printf '// +build tools\npackage tools\nimport _ "github.com/99designs/gqlgen"\nimport _ "github.com/99designs/gqlgen/graphql/introspection"\n' | gofmt > tools.go

    # install gqlgen library
	go mod tidy
	go mod vendor
	go run "github.com/99designs/gqlgen" generate

    # cleanup
	rm -f tools.go
	rm -rf internal/graph/generate/generated
	go mod tidy
	go mod vendor

    # go env -w GOFLAGS=-mod=mod \
    # go run github.com/99designs/gqlgen generate

kill-port:
    lsof -i tcp:4001 \
    kill -9 10551
