OWNER := cshimegi
IMAGE := blog
SEPARATOR := "====================================================================="
GO_DOCKERFILE := ./api/Dockerfile
DB_DOCKERFILE := ./infra/db/Dockerfile
SERVICES := users posts

# Variables
PLATFORM ?= linux/x86_64
VERSION ?= latest
PORT ?= 8080
GO_VERSION ?=
GOV ?=


define check_var
	if [ -z $(call $(1)) ]; then \
		echo "[Error]:$(1) is required"; \
		exit 1; \
	fi
endef

define check_required_vars
	$(call check_var, GOV)
endef

define now
	$(shell date +"%Y/%m/%d %T")" [Execute] ===> "
endef


# help
.DEFAULT_GOAL: help
.PHONY: help
help:
	@ echo "Usage: make [target]"
	@ echo "Targets:"
	@ echo "  build-mono GO_VERSION=<required-go-version> PORT=<optional-port>"
	@ echo "  help"


# build image
#.PHONY: build build-db $(SERVICES)

#build-db:
#	@ echo $(SEPARATOR)
#	@ echo $(SEPARATOR)
#	@ echo $(call now) "Start building db image"
#	@ echo $(SEPARATOR)
#	docker build --platform $(PLATFORM) -f $(DB_DOCKERFILE) -t $(OWNER)/$(IMAGE)-db:$(VERSION) .
#	@ echo $(SEPARATOR)
#	@ echo $(call now) "Done building service: [db]"
#	@ echo $(SEPARATOR)
#
#build: build-db $(SERVICES)
#$(SERVICES):
#	@ echo $(SEPARATOR)
#	@ echo $(call now) "Start building $@ image"
#	@ echo $(SEPARATOR)
#	docker build --build-arg="GO_VERSION=$(GO_VERSION)" \
#				--build-arg="SERVICE_NAME=$(@)" \
#				--build-arg="PORT=$(PORT)" \
#				--platform $(PLATFORM) \
#				-f $(GO_DOCKERFILE) \
#				-t $(OWNER)/$(IMAGE)-$@:$(VERSION) .
#	@ echo $(SEPARATOR)
#	@ echo $(call now) "Done building service: [$@]"
#	@ echo $(SEPARATOR)


build-mono:
	@ echo $(SEPARATOR)
	@ echo $(call now) "Start building mono image"
	@ echo $(SEPARATOR)
	docker build --build-arg="GO_VERSION=$(GO_VERSION)" \
				--build-arg="PORT=$(PORT)" \
				--platform $(PLATFORM) \
				-f $(GO_DOCKERFILE) \
				-t $(OWNER)/$(IMAGE)-mono:$(VERSION) .
	@ echo $(SEPARATOR)
	@ echo $(call now) "Done building mono service"
	@ echo $(SEPARATOR)
