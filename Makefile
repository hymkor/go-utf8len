build:
	go fmt
	go build

test:
	go test -v

bench:
	go test -bench . -benchmem
