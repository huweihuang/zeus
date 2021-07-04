VERSION?=latest

.PHONY: all clean test build

all: fmt vet build image

build: clean
	bash hack/build.sh

image: build
	bash hack/image.sh $(VERSION)

test: fmt vet
	go test -v ./pkg/... ./cmd/... -coverprofile cover.out

fmt:
	go fmt ./pkg/... ./cmd/...

vet:
	go vet ./pkg/... ./cmd/...

clean:
	-rm -Rf _output
