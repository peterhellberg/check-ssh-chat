BINARY = check-ssh-chat
TEST_HOST = pi.c7.se

all: $(BINARY)

**/*.go:
	go build ./...

$(BINARY): **/*.go *.go
	go build -ldflags "-X main.buildCommit `git rev-parse --short HEAD`" .

deps:
	go get .

build: $(BINARY)

clean:
	rm -f $(BINARY)

test: $(BINARY)
	./$(BINARY) -h $(TEST_HOST) -v
