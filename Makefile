CURRENT_DIRECTORY := $(shell pwd)
TESTS_TO_RUN := $(shell go list ./... | grep -v /integrationTests/ | grep -v mock)

build:
	go build ./...

build-cmd:
	(cd cmd && go build)

clean-test:
	go clean -testcache

test: clean-test
	go test ./...

test-coverage:
	@echo "Running unit tests"
	CURRENT_DIRECTORY=$(CURRENT_DIRECTORY) go test -cover -coverprofile=coverage.txt -covermode=atomic -v ${TESTS_TO_RUN}

slow-tests: clean-test
	@docker compose -f docker/docker-compose.yml build
	@docker compose -f docker/docker-compose.yml up -d
	@echo "Waiting for Sui network ports 9000 and 9123 to become available..."
	@bash -c 'for i in $$(seq 1 60); do if (echo >/dev/tcp/localhost/9000) >/dev/null 2>&1 && (echo >/dev/tcp/localhost/9123) >/dev/null 2>&1; then exit 0; fi; sleep 1; done; echo "Sui network did not become ready in time" >&2; exit 1'
	@go test ./integrationTests/... -v -timeout 40m -tags slow
	@docker compose -f docker/docker-compose.yml down -v

lint-install:
ifeq (,$(wildcard test -f bin/golangci-lint))
	@echo "Installing golint"
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s
endif

run-lint:
	@echo "Running golint"
	bin/golangci-lint run --max-issues-per-linter 0 --max-same-issues 0 --timeout=2m

lint: lint-install run-lint