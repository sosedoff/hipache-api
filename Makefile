build:
	go build

release:
	gox -osarch="darwin/amd64 linux/amd64" -output="./bin/hipache-api_{{.OS}}_{{.Arch}}"

deps:
	go get

clean:
	rm -f hipache-api
	rm -f bin/*