## Makefile for the wso2 scripts found in this repo

WSO2SCRIPTS_BUILD_DIR = target
export WSO2SCRIPTS_BUILD_DIR

GOOS=$(shell uname | tr '[A-Z]' '[a-z]')
ifeq ($(shell uname -m),x86_64)
    GOARCH=amd64
else
    GOARCH=386
endif

export GOOS
export GOARCH

all: setup
	hack/build.sh

cross: setup
	hack/crossBuild.sh

setup:
	mkdir -p $(WSO2SCRIPTS_BUILD_DIR)

clean:
	rm -rf $(WSO2SCRIPTS_BUILD_DIR)
