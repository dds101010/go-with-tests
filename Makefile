test:
	go test ./... -count=1 -cover
build:
	go build ./...
fmt:
	go fmt ./...
commit:
	git add .
	git commit
lint:
	errcheck ./...
