.PHONY: start
start:
	present

.PHONY: setup
setup:
	go install golang.org/x/tools/cmd/present@latest
