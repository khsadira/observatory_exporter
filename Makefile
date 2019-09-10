NAME	:=	observatory_exporter
GO		:=	go
FMT		=	gofmt
pkgs	=	$(shell env GO111MODULE=on $(GO) list -m)

FILE	=	main.go\
			metrics_page.go\
			collector_tls.go\
			collector_http.go\

DOCKER_IMAGE_NAME       ?= observatory_exporter

all: format build

test:
	@echo ">> running tests"
	@go test -short $(pkgs)

format:
	@echo ">> formatting code"
	@$(FMT) -w $(FILE)

module:
	@echo ">> creating module"
	@$(GO) mod init $(NAME)

build: 
	@echo ">> building binaries"
	@$(GO) build -o $(NAME)

docker: all
	@echo ">> building docker image"
	@docker build -t $(DOCKER_IMAGE_NAME) .

fclean:
	rm -rf $(NAME) go.sum go.mod

re: fclean module all test

.PHONY: all format build test
