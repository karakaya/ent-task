FROM --platform=$BUILDPLATFORM golang:1.23.3 AS builder
#had to use --platform as i got an err on my M2 mac

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=on

ARG TARGETOS
ARG TARGETARCH

ENV GOOS=$TARGETOS \
    GOARCH=$TARGETARCH

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -trimpath -o api ./cmd/api-service

#user distroless to occupy min space
FROM --platform=$TARGETPLATFORM gcr.io/distroless/static:latest

USER nonroot:nonroot

WORKDIR /app

COPY --from=builder /app/api .
COPY --from=builder /app/migrations ./migrations

ENTRYPOINT ["/app/api"]