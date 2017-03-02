SOURCEDIR=src
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

BINARY=brainfuck
BINDIR=bin

VERSION=1.0.0
BUILD_TIME=`date +%FT%T%z`

.DEFAULT_GOAL: $(BINARY)

$(BINARY): $(SOURCES)
	mkdir -p $(BINDIR)
	go build -o $(BINDIR)/${BINARY} $(SOURCEDIR)/main.go

.PHONY: install
install:
	go install ./...

.PHONY: clean
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
