FROM public.ecr.aws/lambda/go:latest as builder

# Copy the Go code and build the binary
WORKDIR /var/task
COPY --link . .
RUN go build -o whisperer

# Now use a fresh image to reduce the final image size
FROM public.ecr.aws/lambda/go:latest
WORKDIR /var/task
COPY --from=builder /var/task/whisperer .
CMD ["./whisperer"]
