COMMIT := $(shell jj log --template 'commit_id.short(8)' --no-graph --limit 1)
LDFLAGS := -X 'main.version=$(COMMIT)'

.PHONY: build linux clean

build:
	go build -ldflags "$(LDFLAGS)"

linux:
	GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)"

clean:
	rm -f ctags-lsp
