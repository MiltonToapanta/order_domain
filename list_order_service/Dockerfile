# Compilation phase
FROM golang:1.23.4-alpine as builder

WORKDIR /app

# Copy go.mod and go.sum (if available)
COPY go.mod ./
RUN go mod tidy  # Esto genera go.sum automáticamente si no está presente

# Copy the entire project
COPY . .

# Compile the binary
RUN go build -o main .

# If .env exists, it is copied to env.file; otherwise, an empty file is created.
RUN if [ -f .env ]; then cp .env env.file; else touch env.file; fi

# Final phase
FROM alpine:latest

# Install required dependencies (mysql-client for MySQL, not postgresql-client)
RUN apk --no-cache add mysql-client

WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main ./

# Copy the .env file from the builder stage (renamed to .env in the final image)
COPY --from=builder /app/env.file .env

EXPOSE 4002

CMD ["./main"]
