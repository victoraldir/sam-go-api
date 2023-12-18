STACK_NAME ?= bday-api
FUNCTIONS := put-birthday get-birthday authorizer
REGION := us-east-1

# To try different version of Go
GO := go

# Make sure to install aarch64 GCC compilers if you want to compile with GCC.
CC := aarch64-linux-gnu-gcc
GCCGO := aarch64-linux-gnu-gccgo-10

build:
	${MAKE} ${MAKEOPTS} $(foreach function,${FUNCTIONS}, build-${function})

build-%:
	cd app/functions/$* && GOOS=linux GOARCH=arm64 CGO_ENABLED=0 ${GO} build -o bootstrap

deploy: build
	if [ -f samconfig.toml ]; \
		then sam deploy --stack-name ${STACK_NAME} --region ${REGION} --no-confirm-changeset; \
		else sam deploy -g --stack-name ${STACK_NAME} --region ${REGION} --no-confirm-changeset; \
  	fi

test:
	@cd app && go test -tags=unit -race -coverprofile=../coverage.txt -covermode=atomic ./...

run-local:
	@docker-compose up -d
	@sam local start-api --docker-network sam-app-network -n environments/local.json
	
delete:
	@sam delete --stack-name ${STACK_NAME} --region ${REGION}

clean:
	@rm $(foreach function,${FUNCTIONS}, app/functions/${function}/bootstrap)

export GOBIN ?= $(shell pwd)/bin

STATICCHECK = $(GOBIN)/staticcheck

# Many Go tools take file globs or directories as arguments instead of packages
GO_FILES := $(shell \
	       find . '(' -path '*/.*' -o -path './vendor' ')' -prune \
	       -o -name '*.go' -print | cut -b3-)
MODULE_DIRS = app

.PHONY: lint
lint: $(STATICCHECK)
	@rm -rf lint.log
	@echo "Checking formatting..."
	@gofmt -d -s $(GO_FILES) 2>&1 | tee lint.log
	@echo "Checking vet..."
	@$(foreach dir,$(MODULE_DIRS),(cd $(dir) && go vet ./... 2>&1) &&) true | tee -a lint.log
	@echo "Checking staticcheck..."
	@$(foreach dir,$(MODULE_DIRS),(cd $(dir) && $(STATICCHECK) ./... 2>&1) &&) true | tee -a lint.log
	@echo "Checking for unresolved FIXMEs..."
	@git grep -i fixme | grep -v -e Makefile | tee -a lint.log
	@[ ! -s lint.log ]
	@rm lint.log
	@echo "Checking 'go mod tidy'..."
	@make tidy
	@if ! git diff --quiet; then \
		echo "'go diff tidy' resulted in chnges or working tree is dirty:"; \
		git --no-pager diff; \
	fi

$(STATICCHECK):
	cd tools && go install honnef.co/go/tools/cmd/staticcheck

.PHONY: tidy
tidy:
	@$(foreach dir,$(MODULE_DIRS),(cd $(dir) && go mod tidy) &&) true
