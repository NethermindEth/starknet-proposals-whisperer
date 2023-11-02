build-linux:
	@cp build/Dockerfile .
	@docker build -t go-builder .
	@rm Dockerfile
	@mkdir -p bin
	@docker run --rm -v "$(PWD)/bin:/go/src/github.com/NethermindEth/starknet-proposals-whisperer/bin" go-builder
	@mkdir -p compressed
	@make compress

compress:
	@rm -f compressed/whisperer.zip
	@zip compressed/whisperer ./bin/whisperer

install-act:
	@curl -s https://raw.githubusercontent.com/nektos/act/master/install.sh | sudo bash

test-ci:
	@bin/act -P ubuntu-latest=nektos/act-environments-ubuntu:18.04 --secret-file my.secrets --var-file my.variables