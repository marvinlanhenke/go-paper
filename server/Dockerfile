FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o go-paper ./cmd/server/

FROM alpine:latest AS runner

WORKDIR /app

COPY --from=builder /app/go-paper .

ENTRYPOINT [ "./go-paper" ]
