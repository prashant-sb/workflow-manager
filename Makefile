BINARY=workflow-mgr
BIN_DIR=bin
MAIN=./cmd/main.go

# linter versions
#GOLANGCI_VERSION = v1.62.2
GOLANGCI_VERSION = v2.4.0
YAMLLINT_VERSION = 1.35.1

GOLANGCI_LINT = $(BIN_DIR)/golangci-lint
YAMLLINT = $(BIN_DIR)/yamllint

.PHONY: all lint build clean run test deps

all: deps lint build

deps: $(GOLANGCI_LINT) $(YAMLLINT)

$(GOLANGCI_LINT):
	@echo "==> Installing golangci-lint locally..."
	@mkdir -p $(BIN_DIR)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
		| sh -s -- -b $(BIN_DIR) $(GOLANGCI_VERSION)

$(YAMLLINT):
	@echo "==> Installing yamllint locally..."
	@mkdir -p $(BIN_DIR)
	python3 -m venv $(BIN_DIR)/venv
	$(BIN_DIR)/venv/bin/pip install yamllint==$(YAMLLINT_VERSION)
	ln -sf $(PWD)/$(BIN_DIR)/venv/bin/yamllint $(YAMLLINT)

lint: lint-go lint-yaml

lint-go: $(GOLANGCI_LINT)
	@echo "==> Running golangci-lint..."
	$(GOLANGCI_LINT) run ./...

lint-yaml: $(YAMLLINT)
	@echo "==> Linting YAML files..."
	$(YAMLLINT) pkg/dagdef/taskconfig.yaml

build: $(MAIN)
	@echo "==> Building binary..."
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BINARY) $(MAIN)

test:
	@echo "==> Running tests..."
	go test ./...

run: build
	@echo "==> Running..."
	./$(BIN_DIR)/$(BINARY)

clean:
	@echo "==> Cleaning..."
	rm -rf $(BIN_DIR)