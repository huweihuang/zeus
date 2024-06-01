VERSION?=latest

.PHONY: all
all: fmt vet build

##@ General
.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)


##@ Development
.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: test
test: fmt vet ## Run tests.
	go test -v ./pkg/... ./cmd/... -coverprofile cover.out

.PHONY: lint
lint: golangci-lint ## Run golangci-lint linter & yamllint
	$(GOLANGCI_LINT) run

.PHONY: lint-fix
lint-fix: golangci-lint ## Run golangci-lint linter and perform fixes
	$(GOLANGCI_LINT) run --fix


##@ Build
.PHONY: build
build: fmt vet ## Build binary.
	bash hack/build.sh

.PHONY: image
image: build ## Build image.
	bash hack/image.sh $(VERSION)

.PHONY: run
run: build ## Run a server from your host.
	bin/zeus -c configs/config.yaml

.PHONY: clean
clean: ## Clean binary
	-rm -Rf bin

.PHONY: swag
swag: ## generate swagger json
	swag fmt
	swag init -g ./pkg/server/router.go

##@ Build Dependencies

install-swag: ## install swag
	go install github.com/swaggo/swag/cmd/swag@latest

GOLANGCI_LINT = $(shell pwd)/bin/golangci-lint
GOLANGCI_LINT_VERSION ?= v1.54.2
golangci-lint:
	@[ -f $(GOLANGCI_LINT) ] || { \
	set -e ;\
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell dirname $(GOLANGCI_LINT)) $(GOLANGCI_LINT_VERSION) ;\
	}
