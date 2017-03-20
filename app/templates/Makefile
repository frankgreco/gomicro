SOURCEDIR=.
SOURCES := $(shell find . -name '*.go')

.DEFAULT_GOAL: build

build: $(SOURCES)
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
		curl -o ca-certificates.crt https://raw.githubusercontent.com/bagder/ca-bundle/master/ca-bundle.crt

.PHONY: install
install:
		glide install

.PHONY: clean
clean:
		if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
