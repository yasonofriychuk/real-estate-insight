FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN go build -o service ./cmd/service/main.go
RUN go build -o cron ./cmd/cron/import_osm/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /build/service ./service
COPY --from=builder /build/cron ./cron

CMD ["/app/service"]
