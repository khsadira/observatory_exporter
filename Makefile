GO		:=	go
FMT		=	gofmt
pkgs	=	$(shell env GO111MODULE=on $(GO) list -m)

FILE	=	main.go\
			metricsPage.go\
			collector.go\

DOCKER_IMAGE_NAME       ?= exporter_observatory

all: format build

test:
	@echo ">> running tests"
	@go test -short $(pkgs)

format:
	@echo ">> formatting code"
	@$(FMT) -w $(FILE)

module:
	@echo ">> creating module"
	@$(GO) mod init exporter_observatory

build: 
	@echo ">> building binaries"
	@$(GO) build

docker: all
	@echo ">> building docker image"
	@docker build -t $(DOCKER_IMAGE_NAME) .

fclean:
	rm -rf exporter_observatory go.sum go.mod

re: fclean module all test

.PHONY: all format build test
