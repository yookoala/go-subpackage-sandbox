all: clean run test

go-mod:
	@echo -e "\e[1;4;37mGo module operations\e[0m"
	go mod init github.com/yookoala/go-subpackage-sandbox/examples/app
	go mod tidy
	go get -u github.com/yookoala/go-subpackage-sandbox/subpackage
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
	@echo -e "\e[1;4;37mCheck to make sure the parent package and subpackage are both referenced in generated go.mod\e[1;0m"
	@grep -q 'github.com/yookoala/go-subpackage-sandbox' go.mod && echo -e "parent package: \e[32mfound\e[0m" || exit 1
	@grep -q 'github.com/yookoala/go-subpackage-sandbox/subpackage' go.mod && echo -e "subpackage: \e[32mfound\e[0m" || exit 1
	@grep -q 'github.com/goccy/go-json' go.mod && echo -e "subpackage dependency: \e[32mfound\e[0m" || exit 1
	@echo -e "\e[32msuccess\e[0m"

clean:
	@echo -e "\e[1;4;37mClean previous artifact, if any\e[1;0m"
	rm -f go.mod
	rm -f go.sum
	rm -f ./app
	@echo

.PHONY: clean go-mod run test
