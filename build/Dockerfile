FROM amazonlinux:2

# Install Go
RUN amazon-linux-extras enable golang1.11 && yum -y install golang

# Set up the build environment
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

# Copy the project files into the container
WORKDIR $GOPATH/src/github.com/NethermindEth/starknet-proposals-whisperer
COPY . .

# Build the Go binary
RUN GOOS=linux GOARCH=amd64 go build -o ./bin/whisperer
