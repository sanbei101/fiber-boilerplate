ARG TARGETOS
ARG TARGETARCH

FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum /app
RUN go mod download
COPY . .
RUN apk update && apk add upx ca-certificates openssl && update-ca-certificates
RUN echo "Building for OS: ${TARGETOS}, Arch: ${TARGETARCH}, Platform: ${TARGETPLATFORM}" && \
    CGO_ENABLED=0 \
    GOOS=${TARGETOS} \
    GOARCH=${TARGETARCH} \
    go build -ldflags="-s -w" -o /bin/api-binary ./cmd/main.go
RUN upx --best --lzma /bin/api-binary

FROM gcr.io/distroless/static:nonroot
WORKDIR /app/
COPY --from=builder /bin/api-binary /bin/api-binary
COPY --from=builder --chown=nonroot /app/config /app/config
ENTRYPOINT ["/bin/api-binary"]