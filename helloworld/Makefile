NAME=myhelloworld
VERSION=0
RELEASE=0
GOVERSION=$(shell go version|cut -f3 -d' ')

# Build Variables
GOCONTAINER=golang:1.8.1
SRCPATH=mydocker/hellowork

# Container Variables
NAMESPACE=main

# Terminal Color
COLOR_NORMAL=""
COLOR_RESET="\033[m"
COLOR_BOLD="\033[1m"
COLOR_RED="\033[31m"
COLOR_GREEN="\033[32m"
COLOR_YELLOW="\033[33m"


WDIR=$(shell pwd)

default: build container

build: clean binaries

container:
	@echo ${COLOR_BOLD}➭${COLOR_RESET} ${COLOR_GREEN}Building ${NAME} Version ${VERSION}-${RELEASE}${COLOR_RESET}
	@docker build --rm -q -t ${NAME}:${VERSION}-${RELEASE}  .
	@echo ${COLOR_BOLD}${NAME}:${VERSION}-${RELEASE}${COLOR_RESET}

# Install golang dependencies, it should be used as a requirement for other targets
# see target ${NAME}
dependencies:
	@echo ${COLOR_BOLD}➭${COLOR_RESET} ${COLOR_GREEN}Fetching Dependencies with govendor${COLOR_RESET}
	@go get -u github.com/kardianos/govendor
	@govendor fetch +missing
	@echo ${COLOR_GREEN} ..done${COLOR_RESET}

# Starts a docker containter and then calls `make ${NAME}` this fetched the dependencies
# and then builds the binary inside the containet
$(NAME):
	@echo ${COLOR_BOLD}➭${COLOR_RESET} ${COLOR_GREEN}Building ${COLOR_RED}${NAME}${COLOR_RESET} ${COLOR_GREEN}statically${COLOR_RESET}
	@docker run --network=host --rm -v `pwd`:/go/src/${SRCPATH}\
                -w /go/src/${SRCPATH} ${GOCONTAINER}\
                /bin/sh -c "go build -a -o  .build/${NAME} ."
	@echo ${COLOR_GREEN} ..done${COLOR_RESET}


binaries: $(NAME)

all:
	@docker run --network=host --rm -v `pwd`:/go/src/${SRCPATH}\
                -w /go/src/${SRCPATH}/ ${GOCONTAINER}\
                make binaries

clean:
	@echo ${COLOR_BOLD}➭${COLOR_RESET} ${COLOR_GREEN}Cleaning artifacts${COLOR_RESET}
	@rm -rf .build || true

