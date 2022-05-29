OWNER := cshimegi
IMAGE := blog
VERSION := latest
PLATFORM := linux/amd64
SEPARATOR := "====================================================================="
PROCESS := $(shell date +"%Y/%m/%d %T")" [Execute] ===> "
GO_VERSION := 1.17
GO_DOCKERFILE := ./infra/go/Dockerfile
SERVICES := users posts

# help
.DEFAULT_GOAL: help
.PHONY: help
help:
	@echo "Build service images : make build"


# build image
.PHONY: build $(SERVICES)
build: $(SERVICES)
$(SERVICES):
	@echo $(SEPARATOR)
	@echo $(PROCESS) "Start building $@ image"
	@echo $(SEPARATOR)
	docker build --build-arg GO_VERSION=$(GO_VERSION) --build-arg SERVICE_NAME=$@ -f $(GO_DOCKERFILE) --platform=$(PLATFORM) -t $(OWNER)/$(IMAGE)-$@:$(VERSION) .
	@echo $(SEPARATOR)
	@echo $(PROCESS) "Done building service: [$@]"
	@echo $(SEPARATOR)


