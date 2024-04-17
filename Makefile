OWNER := cshimegi
IMAGE := restaurant
SEPARATOR := "====================================================================="
GO_DOCKERFILE := ./api/Dockerfile
DB_DOCKERFILE := ./infra/db/Dockerfile
SERVICES := users appetizers

# Variables
VERSION ?= latest
PORT ?= 8080
GO_VERSION ?=


define check_var
	if [ -z $(call $(1)) ]; then \
		echo "[Error]:$(1) is required"; \
		exit 1; \
	fi
endef

define check_required_vars
	$(call check_var, GO_VERSION)
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
	@ echo "  build-api GO_VERSION=<required-go-version> PORT=<optional-port>"
	@ echo "  help"


# build image
#.PHONY: build build-db $(SERVICES)

#build-db:
#	@ echo $(SEPARATOR)
#	@ echo $(SEPARATOR)
#	@ echo $(call now) "Start building db image"
#	@ echo $(SEPARATOR)
#	docker build -f $(DB_DOCKERFILE) -t $(OWNER)/$(IMAGE)-db:$(VERSION) .
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
#				-f $(GO_DOCKERFILE) \
#				-t $(OWNER)/$(IMAGE)-$@:$(VERSION) .
#	@ echo $(SEPARATOR)
#	@ echo $(call now) "Done building service: [$@]"
#	@ echo $(SEPARATOR)


build-api:
	@ $(call check_required_vars)
	@ echo $(SEPARATOR)
	@ echo $(call now) "Start building api image"
	@ echo $(SEPARATOR)
	docker build --build-arg="GO_VERSION=$(GO_VERSION)" \
				--build-arg="PORT=$(PORT)" \
				-f $(GO_DOCKERFILE) \
				-t $(OWNER)/$(IMAGE)-api:$(VERSION) .
	@ echo $(SEPARATOR)
	@ echo $(call now) "Done building api service"
	@ echo $(SEPARATOR)
