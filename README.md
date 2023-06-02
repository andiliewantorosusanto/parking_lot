# Readme
## Guide to run functional test
### Prerequisite
1. Make sure you run installed go version >= `1.15.3`
2. Run your program to serve port `8080`.

### Run Functional Test
1. `cd parking-lot-golang`
2. `go mod download`
3. To run first case: `go run src/test/functional_testing.go -url=http://localhost:8080 -case=1`
4. To run second case: `go run src/test/functional_testing.go -url=http://localhost:8080 -case=2`

### Run the program
1. `go mod download`
2. `go run src/main.go`

### Makefile command
1. local : `run the application on local`
2. build : `build binary executeable`
3. unit-test : `run unit test`
4. func-test-1 : `run functional test case 1`
5. func-test-2 : `run functional test case 2`

### Executable
1. bin/functional_test : `run functional test`
2. bin/parking_lot : `run the application on local`
3. bin/setup : `download package & run unit test`