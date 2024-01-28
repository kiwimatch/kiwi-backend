# syntax=docker/dockerfile:1
# go project multistage build

# Use standard debian based image to build with stability
FROM golang:1.21 AS build

WORKDIR /app

# Copy only files necessary for installing deps for better layer caching
COPY go.* ./

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

COPY . .

# Run with static build (faster build time)
RUN go build \
    -ldflags="-linkmode external -extldflags -static" \
    -tags netgo \
    -o app .

# Use scratch image for production deployment (minize image size)
FROM alpine:latest
WORKDIR /app

# Set gin mode
ENV GIN_MODE=release

RUN apk update && apk add --no-cache curl

COPY --from=build app/app .
COPY .env .env

# default endpoint used by GIN
EXPOSE 8080

CMD ["./app"]