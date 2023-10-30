build-linux:
	env GOOS=linux GOARCH=amd64 go build -o ./bin/whisperer
	mkdir -p compressed
	@make compress

compress:
	@rm -f ./compressed/whisperer.zip
	@zip ./compressed/whisperer ./bin/whisperer
