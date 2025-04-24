all: clean run test

go-mod:
	@echo -e "\e[1;4;37mGo module operations\e[0m"
	go mod init github.com/yookoala/go-subpackage-sandbox/examples/app2
	go mod tidy
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

clean:
	@echo -e "\e[1;4;37mClean previous artifact, if any\e[1;0m"
	rm -f go.mod
	rm -f go.sum
	rm -f ./app
	@echo

.PHONY: clean go-mod run test
