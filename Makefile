OWNER := cshimegi
IMAGE := blog
VERSION := latest
PLATFORM := linux/amd64
SEPARATOR := "====================================================================="
PROCESS := $(shell date +"%Y/%m/%d %T")" [Execute] ===> "
GO_VERSION := 1.17
GO_DOCKERFILE := ./infra/go/Dockerfile
DB_DOCKERFILE := ./infra/db/Dockerfile
SERVICES := users posts

# help
.DEFAULT_GOAL: help
.PHONY: help
help:
	@echo "Build service images : make build"


# build image
.PHONY: build build-db $(SERVICES)

build-db:
	@echo $(SEPARATOR)
	@echo $(PROCESS) "Start building db image"
	@echo $(SEPARATOR)
	docker build --build-arg GO_VERSION=$(GO_VERSION) -f $(DB_DOCKERFILE) --platform=$(PLATFORM) -t $(OWNER)/$(IMAGE)-db:$(VERSION) .
	@echo $(SEPARATOR)
	@echo $(PROCESS) "Done building service: [db]"
	@echo $(SEPARATOR)

build: build-db $(SERVICES)
$(SERVICES):
	@echo $(SEPARATOR)
	@echo $(PROCESS) "Start building $@ image"
	@echo $(SEPARATOR)
	docker build --build-arg GO_VERSION=$(GO_VERSION) --build-arg SERVICE_NAME=$@ -f $(GO_DOCKERFILE) --platform=$(PLATFORM) -t $(OWNER)/$(IMAGE)-$@:$(VERSION) .
	@echo $(SEPARATOR)
	@echo $(PROCESS) "Done building service: [$@]"
	@echo $(SEPARATOR)
