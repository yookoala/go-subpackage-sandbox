# General test of example applications
test: test-app1 test-app2

test-app1:
	cd examples/app1 && echo && make
	@echo

test-app2:
	cd examples/app2 && echo && make
	@echo

.PHONY: test test-app1 test-app2
