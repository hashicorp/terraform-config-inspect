BINARY_NAME=terraform-config-inspect
DIR=./tfconfig/test-fixtures/basics

GOCMD=env GO111MODULE=on go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOINSTALL=$(GOCMD) install
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	$(GOTEST) -v ./...

run: build
	./$(BINARY_NAME) $(DIR)

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

install:
	$(GOINSTALL) -v .

upgrade:
	$(GOGET) -u

prune:
	$(GOMOD) tidy
