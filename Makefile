local:
	go mod download
	go run src/main.go
build:
	go build src/main.go
unit-test:
	cd src/entity
	go test -v
func-test-1:
	go run src/test/functional_testing.go -url=http://localhost:8080 -case=1
func-test-2:
	go run src/test/functional_testing.go -url=http://localhost:8080 -case=2