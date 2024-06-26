FROM golang:1.21.6-alpine3.19 as builder

RUN apk update && apk add --no-cache git make ca-certificates tzdata openssh
WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go mod verify

RUN go build -ldflags "-s -w" -o migrate ./cmd/migrate/main.go

FROM alpine:3.17

RUN apk --no-cache add ca-certificates tzdata ffmpeg && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
RUN adduser -D -g '' appuser

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/migrate /migrate
COPY --from=builder /build/migrations /migrations

USER appuser

CMD ["/migrate"]
