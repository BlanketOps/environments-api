## ------------------------------------------------------------
## BlanketOps Environments – API / Contract Makefile
##
## This repository defines Kubernetes API contracts only.
## It does NOT build images, run controllers, or deploy anything.
## ------------------------------------------------------------

SHELL := /usr/bin/env bash
.SHELLFLAGS := -ec

##@ General

.PHONY: help
help: ## Show available make targets
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} \
	/^[a-zA-Z_-]+:.*##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 }' \
	$(MAKEFILE_LIST)

##@ Tooling

LOCALBIN ?= $(shell pwd)/bin
CONTROLLER_GEN ?= $(LOCALBIN)/controller-gen

CONTROLLER_TOOLS_VERSION ?= v0.20.0

$(LOCALBIN):
	mkdir -p "$(LOCALBIN)"

.PHONY: controller-gen
controller-gen: $(CONTROLLER_GEN) ## Install controller-gen locally if needed

$(CONTROLLER_GEN): $(LOCALBIN)
	@echo "Installing controller-gen $(CONTROLLER_TOOLS_VERSION)"
	GOBIN=$(LOCALBIN) go install sigs.k8s.io/controller-tools/cmd/controller-gen@$(CONTROLLER_TOOLS_VERSION)

##@ Code Generation

.PHONY: generate
generate: controller-gen ## Generate deepcopy code
	$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..."

.PHONY: manifests
manifests: controller-gen ## Generate CRDs from API types
	$(CONTROLLER_GEN) crd paths="./api/..." output:crd:artifacts:config=config/crd/bases

##@ Quality

.PHONY: fmt
fmt: ## Run go fmt
	go fmt ./...

.PHONY: vet
vet: ## Run go vet
	go vet ./...

##@ Verification

.PHONY: verify
verify: fmt vet generate manifests ## Verify API contracts are consistent
