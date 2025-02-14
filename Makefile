#!/usr/bin/make -f

TAG = `git tag | sort -V | tail -1`
DATE = `date`

.PHONY: build
build: 
	GOWORK=off CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -ldflags "-X 'github.com/fh-x4/littletool/cmd.Version=${TAG}' -X 'github.com/fh-x4/littletool/cmd.Date=${DATE}'"
