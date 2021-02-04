.PHONY: update-pkg-cache

VERSION := $(shell git describe --tags --abbrev=0)

update-pkg-cache:
	GOPROXY=https://proxy.golang.org GO111MODULE=on \
	go get github.com/talon-one/talon_go/v2@$(VERSION)
