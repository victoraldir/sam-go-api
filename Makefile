STACK_NAME ?= bday-api
FUNCTIONS := put-birthday get-birthday
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
		then sam deploy --stack-name ${STACK_NAME} --region ${REGION}; \
		else sam deploy -g --stack-name ${STACK_NAME} --region ${REGION}; \
  	fi

test:
	@cd app && go test ./...

run-local:
	@docker-compose up -d
	@sam local start-api --docker-network sam-app-network -n environments/local.json
	
delete:
	@sam delete --stack-name ${STACK_NAME} --region ${REGION}

clean:
	@rm $(foreach function,${FUNCTIONS}, app/functions/${function}/bootstrap)
