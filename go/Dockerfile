FROM golang:1.14-alpine

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Install certificates
RUN apk --no-cache add \
    git \
    ca-certificates \
    && update-ca-certificates
