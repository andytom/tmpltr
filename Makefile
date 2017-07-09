.DEFAULT_GOAL := compile

# -- Variables --
# All the go packages excluding the vendored libs
PKGS = $(shell go list ./... | grep -v /vendor/)

# The git tag for the version
VERSION = $(shell git describe --exact-match --tags 2>/dev/null)

# Build Flags
LDFLAGS = -ldflags "-X cmd.version.version=${VERSION}"

# -- High level targets --
# We only list these targets in the help. The other targets can still be used
# but it is generally better to call one of these.

.PHONY: help
help: ## Prints this help
ifneq ($(.DEFAULT_GOAL),)
	@echo "Default target: \033[36m$(.DEFAULT_GOAL)\033[0m"
endif
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: compile
compile: clean check ## Checks and builds the app for the local machine arch
	@echo "-- Building --"
	@go build ${LDFLAGS}

.PHONY: check
check: test vet lint ## Runs the tests, go vet and golint

# -- Low level targets --
# These targets are more low level and not included in the help. You can call
# them directly but generally you would use the higher level target.

.PHONY: test
test:
	@echo "-- Running Tests --"
	@go test ${PKGS} -cover

.PHONY: lint
lint:
	@echo "-- Running Lint --"
	@golint -set_exit_status ${PKGS}

.PHONY: vet
vet:
	@echo "-- Running Vet --"
	@go vet ${PKGS}

.PHONY: clean
clean:
	@go clean
