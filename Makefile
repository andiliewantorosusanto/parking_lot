local:
	go mod download
	go run main.go
func-test-1:
	go run functional_testing.go -url=http://localhost:8080 -case=1
func-test-2:
	go run functional_testing.go -url=http://localhost:8080 -case=2