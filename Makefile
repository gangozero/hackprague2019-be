SOURCEDIR=.

PROJECT_NAME=hackprague2019-be
PROJECT_ROOT=`git rev-parse --show-toplevel`
PROJECT_TOOL_PATH=/go/src/github.com/gangozero/$(PROJECT_NAME)

SWAGGER_VERSION=v0.19.0
SWAGGER_CMD=docker run --rm -it -v $(PROJECT_ROOT):$(PROJECT_TOOL_PATH) -w $(PROJECT_TOOL_PATH) quay.io/goswagger/swagger:$(SWAGGER_VERSION)

DOCKER_REPO=$(PROJECT_NAME)
VERSION=0.1
TAG=$(DOCKER_REPO)/backend:$(VERSION)

.DEFAULT_GOAL: validate

.PHONY: install-swagger validate vendor

install-swagger:
	docker pull quay.io/goswagger/swagger:$(SWAGGER_VERSION)

validate:
	$(SWAGGER_CMD) validate ./openapi/swagger.yaml

generate-full:
	$(SWAGGER_CMD) generate server -A $(PROJECT_NAME) -f ./openapi/swagger.yaml --target ./generated

generate:
	$(SWAGGER_CMD) generate server -A $(PROJECT_NAME) -f ./openapi/swagger.yaml --target ./generated --exclude-main

vendor:
	go mod vendor

build:
	go build

docker-build:
	docker build --pull -f ./deployment/Dockerfile -t $(TAG) .

docker-push:
	docker push $(TAG) && docker image rm $(TAG)

docker-run:
	docker rm -f $(PROJECT_NAME) || true && docker run -dt --name $(PROJECT_NAME) --rm -p 8080:8080 $(TAG)