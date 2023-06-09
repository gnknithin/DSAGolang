update-mod-packages:
	go mod tidy
clean-go-test-cache:
	go clean -testcache
run-go-tests: clean-go-test-cache
	go test ./...
code-coverage: clean-go-test-cache
	go test -coverpkg=./... ./...
