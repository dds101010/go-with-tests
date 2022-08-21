test:
	go test ./... -count=1 -cover
test_bench:
	go test -v ./concurrency -count=1 -bench=.
build:
	go build ./...
fmt:
	go fmt ./...
commit:
	git add .
	git commit
lint:
	errcheck ./...
