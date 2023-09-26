FROM golang:1.20-alpine3.18 AS builder

COPY wait-for.sh /wait-for.sh
RUN chmod +x /wait-for.sh

RUN apk --no-cache add git
WORKDIR /build
COPY go.mod /src/go.mod
COPY go.sum /src/go.sum
WORKDIR /src
RUN go mod download
COPY . /src
RUN go build -o /build/api ./cmd/api

FROM alpine:3.18 AS service
COPY --from=builder /wait-for.sh /service-bin/wait-for.sh
RUN chmod +x /service-bin/wait-for.sh

EXPOSE 8080/tcp
COPY --from=builder /build/api /service-bin/api
WORKDIR /db
WORKDIR /

ENV DATABASE_NAME="cie-db" \
    DATABASE_USER="postgres" \
    DATABASE_PASSWORD="postgres" \
    DATABASE_HOST="postgresql" \
    DATABASE_PORT="5432" \
    DATABASE_SSLMODE="disable" \
    DATABASE_LOGMODE="false" \
    GIN_MODE="debug" \
    SERVER_PORT="8080" \
    ENV="development" \
    LOG_OUTPUT="./server.log" \
    LOG_LEVEL="info" \
    CONTEXT_TIMEOUT=2


ENTRYPOINT ["/service-bin/api", ":$SERVER_PORT"]