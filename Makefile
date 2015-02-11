build:
	godep go build

release:
	gox -osarch="darwin/amd64 linux/amd64" -output="./bin/hipache-api_{{.OS}}_{{.Arch}}"

deps:
	godep restore

setup:
	go get github.com/tools/godep
	go get github.com/mitchellh/gox

clean:
	rm -f hipache-api
	rm -f bin/*