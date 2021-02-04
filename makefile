.PHONY: update-pkg-cache

update-pkg-cache:
	GOPROXY=https://proxy.golang.org GO111MODULE=on \
	go get github.com/talon-one/talon_go@$(git describe --tags --abbrev=0)
