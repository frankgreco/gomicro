SOURCEDIR=.
SOURCES := $(shell find . -name '*.go')

BINARY=<%= project %>

.DEFAULT_GOAL: $(BINARY)

$(BINARY): $(SOURCES) install
	go build

.PHONY: install
install:
	glide update
	glide install

.PHONY: local-dev
local-dev: $(BINARY)
	docker-compose up -d

.PHONY: docker
docker: install
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
	curl -o ca-certificates.crt https://raw.githubusercontent.com/bagder/ca-bundle/master/ca-bundle.crt
	docker build -t frankgreco/gohttp:latest .

.PHONY: clean
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
