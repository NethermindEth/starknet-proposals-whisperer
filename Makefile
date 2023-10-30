build-linux:
	@cp build/Dockerfile .
	@docker build -t go-builder .
	@rm Dockerfile
	@docker run --rm -v "$(PWD)/bin:/go/src/github.com/NethermindEth/starknet-proposals-whisperer/bin" go-builder
	@make compress

compress:
	@rm -f ./compressed/whisperer.zip
	@zip ./compressed/whisperer ./bin/whisperer
