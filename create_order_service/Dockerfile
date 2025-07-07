# Compilation phase
FROM golang:1.23.4-alpine as builder
 
WORKDIR /app
 
# Copy dependency files and download modules.
COPY go.mod ./
RUN go mod tidy
 
# Copy the entire project (including .env locally, if it exists)
COPY . .
 
# Compile the binary
RUN go build -o main .
 
# Final phase
FROM alpine:latest
 
# Install required dependencies
RUN apk --no-cache add postgresql-client
 
WORKDIR /root/
 
# Copy the compiled binary
COPY --from=builder /app/main .
 
# Copy the file env.file (we renamed it to .env in the final image)
COPY --from=builder /app/env.file .env
 
EXPOSE 4001
 
CMD ["./main"]