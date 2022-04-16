ifndef $(GOPATH)
	GOPATH = $(shell go env GOPATH)
	export GOPATH
endif

install_air:
	go install github.com/cosmtrek/air@latest

dev:
	$(GOPATH)/bin/air;