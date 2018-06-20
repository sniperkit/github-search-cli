.PHONY: all test clean man glide fast release install

GO15VENDOREXPERIMENT=1

PROG_NAME := "ghs"

all: deps-ensure test build install version

build: deps-ensure
	@go build -ldflags "-X github.com/sniperkit/github-search-cli/pkg.Version=`cat VERSION`" -o ./bin/$(PROG_NAME) ./cmd/$(PROG_NAME)/*.go

version: deps-ensure
	@which $(PROG_NAME)
	@$(PROG_NAME) --version

install: deps-ensure
	@go install -ldflags  "-X github.com/sniperkit/github-search-cli/pkg.Version=`cat VERSION`" ./cmd/$(PROG_NAME)
	@$(PROG_NAME) --version

fast: deps
	@go build -i -ldflags "-X github.com/sniperkit/github-search-cli/pkg.Version=`cat VERSION`-dev" -o ./bin/$(PROG_NAME) ./cmd/$(PROG_NAME)/*.go
	@$(PROG_NAME) --version

deps: deps-create deps-ensure

deps-create:
	@rm -f glide.*
	@rm -f *Gopkg*
	@yes no | glide create

deps-ensure:
	@glide install --strip-vendor

test:
	@go test ./pkg/...

clean:
	@go clean
	@rm -fr ./bin
	@rm -fr ./dist

release: $(PROG_NAME)
	@git tag -a `cat VERSION`
	@git push origin `cat VERSION`

cover:
	@rm -rf *.coverprofile
	@go test -coverprofile=$(PROG_NAME).coverprofile ./pkg/...
	@gover
	@go tool cover -html=$(PROG_NAME).coverprofile ./pkg/...

lint: install-deps-dev
	@errors=$$(gofmt -l .); if [ "$${errors}" != "" ]; then echo "$${errors}"; exit 1; fi
	@errors=$$(glide novendor | xargs -n 1 golint -min_confidence=0.3); if [ "$${errors}" != "" ]; then echo "$${errors}"; exit 1; fi

vet:
	@go vet $$(glide novendor)