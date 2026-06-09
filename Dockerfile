FROM golang:1.26.4 AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-s -w" \
    -o /app \
    ./cmd/server

FROM gcr.io/distroless/static-debian12

COPY --from=builder /app /app

USER nonroot:nonroot

ENTRYPOINT ["/app"]
