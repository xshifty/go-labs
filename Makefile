.PHONY: linux

linux:
	go build  -o ./build/graph cmd/graph/main.go
	go build  -o ./build/xml cmd/xml/main.go

clean:
	rm -rf build

