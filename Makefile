build:
	go build

release:
	gox -osarch="darwin/amd64 linux/amd64" -output="./bin/hipache-api_{{.OS}}_{{.Arch}}"

make clean:
	rm -f hipache-api
	rm -f bin/*