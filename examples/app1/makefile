SHELL:=/bin/bash
GOPATH:=$(shell go env GOPATH)

all: clean run test

go-mod:
	@echo -e "\e[1;4;37mGo module operations\e[0m"
	go mod init github.com/yookoala/go-subpackage-sandbox/examples/app
	go mod tidy
	go get -u github.com/yookoala/go-subpackage-sandbox
	@echo

app: go-mod
	@echo -e "\e[1;4;37mBuild application\e[1;0m"
	go build -o $@
	@echo

run: app
	@echo -e "\e[1;4;37mTest run\e[1;0m"
	@./app
	@echo

test:
	@echo -e "\e[1;4;37mCheck to make sure the subpackage is not referenced in generated go.mod\e[1;0m"
	grep -q 'github.com/yookoala/go-subpackage-sandbox/subpackage' go.mod && exit 1 || echo -e "\e[32msuccess\e[0m"
	@echo
	@echo -e "\e[1;4;37mCheck if GOPATH has the dependency of subpackage by mistake\e[1;0m"
	@if [ -d $(GOPATH)/pkg/mod/github.com/goccy/go-json* ]; then \
		echo -e "\e[31mGOPATH has the dependency of subpackage by mistake\e[0m"; \
		exit 1; \
	else \
		echo -e "\e[32msuccess\e[0m"; \
	fi

clean:
	@echo -e "\e[1;4;37mClean previous artifact, if any\e[1;0m"
	rm -f go.mod
	rm -f go.sum
	rm -f ./app
	@if [ -d $(GOPATH)/pkg/mod/github.com/yookoala/go-subpackage-sandbox ]; then \
		echo "cleaning up $(GOPATH)/pkg/mod/github.com/yookoala/go-subpackage-sandbox"; \
		find $(GOPATH)/pkg/mod/github.com/yookoala/go-subpackage-sandbox -type d -exec chmod u+w {} \; ; \
		rm -Rf $(GOPATH)/pkg/mod/github.com/yookoala/go-subpackage-sandbox ; \
	fi
	@if [ -d $(GOPATH)/pkg/mod/github.com/goccy/go-json* ]; then \
		echo "cleaning up $(GOPATH)/pkg/mod/github.com/goccy/go-json*"; \
		find $(GOPATH)/pkg/mod/github.com/goccy/go-json* -type d -exec chmod u+w {} \; ; \
		rm -Rf $(GOPATH)/pkg/mod/github.com/goccy/go-json* ; \
	fi
	@echo

.PHONY: clean go-mod run test
